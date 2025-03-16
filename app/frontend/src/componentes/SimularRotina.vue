<script setup>
import { ref } from 'vue'
import { SimularRotina } from '../../wailsjs/go/main/App'

const means = ref(Array(4).fill(0))      // 4 campos para means
const stds = ref(Array(4).fill(0))       // 4 campos para stds
const tamanho = ref(0)                   // campo para o tamanho
const resultados = ref([])               // armazena os resultados da simulação
const loading = ref(false)               // indicador de carregamento
const erro = ref("")                     // mensagem de erro

async function executarSimulacao() {
  erro.value = ""
  resultados.value = []

  // Validação dos campos
  if (means.value.some(isNaN) || stds.value.some(isNaN) || isNaN(tamanho.value)) {
    erro.value = "Por favor, preencha todos os campos com valores válidos."
    return
  }
  if (tamanho.value <= 0) {
    erro.value = "O tamanho deve ser maior que zero."
    return
  }

  loading.value = true

  try {
    console.log("Dados enviados ao backend:", {
      means: means.value,
      stds: stds.value,
      tamanho: tamanho.value,
    })

    // Chama a função SimularRotina exposta no backend (o binding é feito na struct App)
    resultados.value = await SimularRotina(
      means.value.map(Number),
      stds.value.map(Number),
      Number(tamanho.value)
    )

    console.log("Resultados recebidos:", resultados.value)
  } catch (e) {
    console.error("Erro:", e)
    erro.value = "Falha na simulação. Verifique o console para mais detalhes."
  } finally {
    loading.value = false
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
        <div v-for="(_, index) in 4" :key="'mean' + index">
          <input
            type="number"
            v-model="means[index]"
            :placeholder="'Mean ' + (index + 1)"
            step="0.01"
            required
          />
        </div>
      </div>

      <!-- Inputs para os Desvios Padrão -->
      <div class="form-group">
        <label>Desvios Padrão (Stds):</label>
        <div v-for="(_, index) in 4" :key="'std' + index">
          <input
            type="number"
            v-model="stds[index]"
            :placeholder="'Std ' + (index + 1)"
            step="0.01"
            required
          />
        </div>
      </div>

      <!-- Input para o Tamanho -->
      <div class="form-group">
        <label>Tamanho:</label>
        <input type="number" v-model="tamanho" required />
      </div>

      <!-- Botão de Submissão -->
      <button type="submit" :disabled="loading">
        {{ loading ? 'Simulando...' : 'Rodar Simulação' }}
      </button>
    </form>

    <!-- Exibição dos Resultados -->
    <div v-if="resultados.length > 0">
      <h2>Resultados ({{ resultados.length }} amostras):</h2>
      <pre>{{ resultados }}</pre>
    </div>

    <!-- Exibição de Mensagem de Erro -->
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
pre {
  background-color: #f5f5f5;
  padding: 10px;
  border-radius: 4px;
  overflow-x: auto;
}
</style>
