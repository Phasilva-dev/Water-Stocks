<template>
  <div class="container">
    <h1>Simulação Estocástica</h1>
    
    <!-- Formulário com 9 campos -->
    <form @submit.prevent="executarSimulacao">
      <!-- Means -->
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

      <!-- Stds -->
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

      <!-- Tamanho -->
      <div class="form-group">
        <label>Tamanho:</label>
        <input type="number" v-model="tamanho" required />
      </div>

      <button type="submit">Rodar Simulação</button>
    </form>

    <!-- Resultados -->
    <div v-if="resultados.length > 0">
      <h2>Resultados ({{ resultados.length }} amostras):</h2>
      <pre>{{ resultados }}</pre>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      means: Array(4).fill(0), // 4 campos para means
      stds: Array(4).fill(0),  // 4 campos para stds
      tamanho: 0,
      resultados: [],
    };
  },
  methods: {
    async executarSimulacao() {
      try {
        // Chama o método Go via Wails
        this.resultados = await window.backend.SimularRotina(
          this.means.map(Number),
          this.stds.map(Number),
          Number(this.tamanho)
        );
      } catch (error) {
        console.error("Erro:", error);
        alert("Falha na simulação");
      }
    },
  },
};
</script>

<style>
/* Estilos similares aos anteriores */
</style>