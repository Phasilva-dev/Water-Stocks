<script setup>
import { ref, nextTick } from "vue";
import Plotly from "plotly.js-dist-min"; // Certifique-se de ter instalado via npm (npm install plotly.js-dist-min)
import { SimularRotina } from "../../wailsjs/go/main/App";

// Estados reativos
const means = ref([0, 0, 0, 0]);  // 4 campos para means
const stds = ref([0, 0, 0, 0]);   // 4 campos para stds
const tamanho = ref(10);          // Tamanho padrão para evitar NaN
const resultados = ref([]);       // Armazena os resultados retornados
const loading = ref(false);       // Indicador de carregamento
const erro = ref("");             // Mensagem de erro

// Função para plotar o histograma usando Plotly.js
function plotHistogram() {
  // Como as amostras são intercaladas (0: acordar, 1: sair, 2: voltar, 3: dormir), separamos cada uma:
  const samplesAcordar = resultados.value.filter((_, i) => i % 4 === 0);
  const samplesSair = resultados.value.filter((_, i) => i % 4 === 1);
  const samplesVoltar = resultados.value.filter((_, i) => i % 4 === 2);
  const samplesDormir = resultados.value.filter((_, i) => i % 4 === 3);

  // Definindo os traços para o histograma
  const traceAcordar = {
    x: samplesAcordar,
    type: "histogram",
    opacity: 0.5,
    name: "Acordar",
    marker: { color: "blue" }
  };
  
  const traceSair = {
    x: samplesSair,
    type: "histogram",
    opacity: 0.5,
    name: "Sair",
    marker: { color: "green" }
  };

  const traceVoltar = {
    x: samplesVoltar,
    type: "histogram",
    opacity: 0.5,
    name: "Voltar",
    marker: { color: "red" }
  };

  const traceDormir = {
    x: samplesDormir,
    type: "histogram",
    opacity: 0.5,
    name: "Dormir",
    marker: { color: "purple" }
  };

  // Conjunto de dados para o gráfico
  const data = [traceAcordar, traceSair, traceVoltar, traceDormir];

  // Configurações do layout
  const layout = {
    title: "Histograma de Amostras",
    barmode: "overlay", // Exibe os histogramas sobrepostos para facilitar a comparação
    xaxis: { title: "Valores" },
    yaxis: { title: "Frequência" }
  };

  // Renderiza o gráfico no elemento com id "histogram"
  Plotly.newPlot("histogram", data, layout);
}

// Função que chama o backend e executa a simulação
async function executarSimulacao() {
  erro.value = "";
  resultados.value = [];

  // Validação dos campos
  if (means.value.some((x) => isNaN(x)) || stds.value.some((x) => isNaN(x)) || isNaN(tamanho.value)) {
    erro.value = "Por favor, preencha todos os campos corretamente.";
    return;
  }
  if (tamanho.value <= 0) {
    erro.value = "O tamanho da amostra deve ser maior que zero.";
    return;
  }

  loading.value = true;

  try {
    console.log("Enviando para o backend:", {
      means: means.value,
      stds: stds.value,
      tamanho: tamanho.value,
    });

    // Chama a função do backend via Wails
    resultados.value = await SimularRotina(
      means.value.map(Number),
      stds.value.map(Number),
      Number(tamanho.value)
    );

    console.log("Resultados recebidos:", resultados.value);

    // Aguarda o DOM atualizar e, em seguida, plota o histograma
    await nextTick();
    plotHistogram();
  } catch (e) {
    console.error("Erro ao executar SimularRotina:", e);
    erro.value = "Falha na simulação. Verifique o console para mais detalhes.";
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div class="container">
    <h1>Simulação Estocástica</h1>
    <form @submit.prevent="executarSimulacao">
      <!-- Inputs para os Means -->
      <div class="form-group">
        <label>Means:</label>
        <div v-for="(mean, index) in means" :key="'mean' + index">
          <input
            type="number"
            v-model.number="means[index]"
            :placeholder="'Mean ' + (index + 1)"
            step="0.01"
            required
          />
        </div>
      </div>

      <!-- Inputs para os Desvios Padrão (Stds) -->
      <div class="form-group">
        <label>Desvios Padrão (Stds):</label>
        <div v-for="(std, index) in stds" :key="'std' + index">
          <input
            type="number"
            v-model.number="stds[index]"
            :placeholder="'Std ' + (index + 1)"
            step="0.01"
            required
          />
        </div>
      </div>

      <!-- Input para o Tamanho -->
      <div class="form-group">
        <label>Tamanho:</label>
        <input type="number" v-model.number="tamanho" required />
      </div>

      <!-- Botão de Submissão -->
      <button type="submit" :disabled="loading">
        {{ loading ? "Simulando..." : "Rodar Simulação" }}
      </button>
    </form>

    <!-- Container para o histograma (substituindo o print das amostras) -->
    <div id="histogram"></div>

    <!-- Exibição de mensagem de erro -->
    <div v-if="erro" class="erro">
      <p>{{ erro }}</p>
    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
}
.form-group {
  margin-bottom: 20px;
}
input {
  width: 100%;
  padding: 8px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
}
button {
  padding: 10px 20px;
  background-color: #42b983;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
button:hover:not(:disabled) {
  background-color: #369f6e;
}
.erro {
  margin-top: 20px;
  padding: 10px;
  background-color: #ffebee;
  border: 1px solid #ffcdd2;
  border-radius: 4px;
  color: #c62828;
}
</style>
