// frontend/src/App.jsx

import { useState } from 'react';
import './App.css';
import SimulationForm from './pages/SimulationForm';
import ResultsViewer from './pages/ResultsViewer';
import AboutPage from './pages/AboutPage';

// Um ícone simples de gota d'água usando SVG
const WaterDropIcon = () => (
  <svg viewBox="0 0 384 512" width="40" height="50" fill="currentColor">
    <path d="M192 512C86 512 0 426 0 320C0 228.8 130.2 57.7 166.6 11.7C172.6 4.2 181.5 0 191.1 0C201.5 0 210.4 4.2 216.4 11.7C252.8 57.7 384 228.8 384 320C384 426 298 512 192 512z"/>
  </svg>
);


function App() {
  // State para controlar qual página está visível
  const [currentPage, setCurrentPage] = useState('home');

  const renderPage = () => {
    switch (currentPage) {
      case 'new_simulation':
        return <SimulationForm />;
      case 'view_results':
        return <ResultsViewer />;
      case 'about':
        return <AboutPage />;
      default:
        // A página inicial com os botões de navegação
        return (
          <div className="home-container">
            <p>Selecione uma opção para começar</p>
            <nav className="main-nav">
              <button className="nav-button" onClick={() => setCurrentPage('new_simulation')}>
                Nova Simulação
              </button>
              <button className="nav-button" onClick={() => setCurrentPage('view_results')}>
                Visualizar Resultados
              </button>
              <button className="nav-button" onClick={() => setCurrentPage('about')}>
                Sobre o Projeto
              </button>
            </nav>
          </div>
        );
    }
  };

  return (
    <div id="App">
      <header className="app-header">
        <WaterDropIcon />
        <h1>HydroSim</h1>
        <h2>Simulador de Demanda Hídrica</h2>
      </header>
      
      <main className="content">
        {/* Renderiza um botão "Voltar" se não estivermos na página inicial */}
        {currentPage !== 'home' && (
          <button className="back-button" onClick={() => setCurrentPage('home')}>
            &larr; Voltar ao Início
          </button>
        )}
        {renderPage()}
      </main>
    </div>
  );
}

export default App;