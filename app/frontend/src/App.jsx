import { useState, useEffect } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
// Importe a sua nova função do backend Go
import { RunSimulation } from "../wailsjs/go/main/App";

function App() {
    // State para exibir mensagens na UI
    const [resultText, setResultText] = useState("Configure os parâmetros e inicie a simulação.");

    // States para cada parâmetro da simulação, com valores padrão
    const [size, setSize] = useState(10000);
    const [day, setDay] = useState(7);
    const [toiletType, setToiletType] = useState(1);
    const [showerType, setShowerType] = useState(1);
    const [filename, setFilename] = useState("simulacao-react");

    // Hook para escutar eventos do Go.
    // Isso é executado apenas uma vez quando o componente é montado.
    useEffect(() => {
        // Escuta pelo evento 'simulationComplete' que definimos no backend
        window.runtime.EventsOn('simulationComplete', (message) => {
            console.log("Evento recebido do Go:", message);
            // Atualiza a UI com a mensagem final da simulação
            setResultText(message);
        });

        // Função de limpeza: remove o listener quando o componente é desmontado
        return () => {
            window.runtime.EventsOff('simulationComplete');
        };
    }, []); // O array vazio [] garante que isso rode apenas uma vez

    // Função que será chamada pelo botão "Iniciar Simulação"
    function startSimulation() {
        // Atualiza a UI para dar feedback imediato ao usuário
        setResultText(`Iniciando simulação '${filename}' com ${size} casas...`);

        // Chama a função do backend Go com os valores atuais do estado
        RunSimulation(size, day, toiletType, showerType, filename)
            .then((immediateResponse) => {
                // A resposta imediata do Go ("Simulação iniciada...") vem aqui
                console.log(immediateResponse);
                setResultText(immediateResponse);
            })
            .catch((err) => {
                // Caso ocorra algum erro ao tentar chamar a função
                console.error("Erro ao chamar RunSimulation:", err);
                setResultText(`Erro: ${err}`);
            });
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <h1>Simulador Hidrológico</h1>
            <div id="result" className="result">{resultText}</div>
            
            {/* Seção com os inputs para os parâmetros da simulação */}
            <div className="simulation-form">
                <div className="input-box">
                    <label htmlFor="filename">Nome do Arquivo:</label>
                    <input id="filename" className="input" value={filename} onChange={(e) => setFilename(e.target.value)} type="text"/>
                </div>
                <div className="input-box">
                    <label htmlFor="size">Número de Casas:</label>
                    <input id="size" className="input" value={size} onChange={(e) => setSize(parseInt(e.target.value))} type="number"/>
                </div>
                <div className="input-box">
                    <label htmlFor="day">Dias de Simulação:</label>
                    <input id="day" className="input" value={day} onChange={(e) => setDay(parseInt(e.target.value))} type="number"/>
                </div>
                <div className="input-box">
                    <label htmlFor="toiletType">Tipo de Vaso (1-4):</label>
                    <input id="toiletType" className="input" value={toiletType} onChange={(e) => setToiletType(parseInt(e.target.value))} type="number"/>
                </div>
                <div className="input-box">
                    <label htmlFor="showerType">Tipo de Chuveiro (1-2):</label>
                    <input id="showerType" className="input" value={showerType} onChange={(e) => setShowerType(parseInt(e.target.value))} type="number"/>
                </div>

                {/* Botão que dispara a simulação */}
                <button className="btn" onClick={startSimulation}>
                    Iniciar Simulação
                </button>
            </div>
        </div>
    )
}

export default App;