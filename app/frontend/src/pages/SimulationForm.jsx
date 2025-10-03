// --- START OF FILE SimulationForm.jsx ---
import { useState, useEffect } from 'react';
// Importe a função real do seu backend Go gerado pelo Wails
import { RunSimulation } from "../../wailsjs/go/main/App";

function SimulationForm() {
    // --- STATE PARA OS DADOS DO FORMULÁRIO ---
    const [params, setParams] = useState({
        filename: "simulacao-teste",
        size: 10000,
        day: 7,
        toiletType: 1,
        showerType: 1,
    });

    // --- STATE PARA VALIDAÇÃO E FEEDBACK ---
    const [errors, setErrors] = useState({});
    const [statusMessage, setStatusMessage] = useState("Configure os parâmetros e inicie a simulação.");
    const [isSimulating, setIsSimulating] = useState(false);
    const [isError, setIsError] = useState(false); // Novo state para controlar a cor da mensagem

    // --- HANDLER PARA ATUALIZAR O ESTADO DOS PARÂMETROS ---
    const handleInputChange = (e) => {
        const { name, value } = e.target;
        // Permite apenas números para campos numéricos, exceto o nome do arquivo
        const processedValue = (e.target.type === 'number' || e.target.tagName === 'SELECT') ? parseInt(value) || 0 : value;
        setParams(prevParams => ({
            ...prevParams,
            [name]: processedValue,
        }));
    };

    // --- FUNÇÃO DE VALIDAÇÃO ---
    const validateParams = () => {
        const newErrors = {};
        if (!params.filename.trim()) newErrors.filename = "Nome do arquivo é obrigatório.";
        if (params.size <= 0) newErrors.size = "Deve ser um valor positivo.";
        if (params.day <= 0) newErrors.day = "Deve ser um valor positivo.";
        if (params.toiletType < 1 || params.toiletType > 4) newErrors.toiletType = "Opção inválida.";
        if (params.showerType < 1 || params.showerType > 2) newErrors.showerType = "Opção inválida.";

        setErrors(newErrors);
        // Retorna true se não houver erros
        return Object.keys(newErrors).length === 0;
    };

    // --- FUNÇÃO PARA INICIAR A SIMULAÇÃO ---
    const startSimulation = () => {
        if (!validateParams()) {
            setStatusMessage("Por favor, corrija os erros no formulário.");
            setIsError(true); // Define como erro para estilização
            return;
        }

        setIsSimulating(true);
        setIsError(false); // Reseta o estado de erro
        setStatusMessage(`Iniciando simulação '${params.filename}'...`);
        
        RunSimulation(params.size, params.day, params.toiletType, params.showerType, params.filename)
            .then(response => {
                // O backend retorna uma mensagem de sucesso
                setStatusMessage(response);
                setIsError(false); // Garante que não é um erro
            })
            .catch(err => {
                // O backend retorna um erro
                setStatusMessage(`Erro: ${err}`);
                setIsError(true); // Marca como erro
            })
            .finally(() => {
                // Isso será executado tanto em caso de sucesso quanto de erro
                setIsSimulating(false); // Libera o botão e os inputs
            });
    };

    // --- EFEITO PARA OUVIR EVENTOS DE STATUS DO BACKEND (WAILS) ---
    useEffect(() => {
        // Este evento é para mostrar o progresso durante a simulação
        const cleanupStatus = window.runtime.EventsOn('simulationStatus', (message) => {
            // Só atualiza a mensagem se a simulação estiver em andamento
            if (isSimulating) {
                setStatusMessage(message);
            }
        });

        // A função de limpeza é chamada quando o componente é desmontado
        return () => {
            cleanupStatus();
        };
    }, [isSimulating]); // A dependência garante que o listener seja gerenciado corretamente

    return (
        <div className="page-container">
            <h3>Criar Nova Simulação</h3>
            <p>Defina os parâmetros abaixo para gerar um novo arquivo de análise.</p>
            
            <div className="simulation-form compact">
                {/* --- INPUTS COM VALIDAÇÃO E PLACEHOLDERS --- */}
                <div className="input-box">
                    <label htmlFor="filename">Nome do Arquivo</label>
                    <input name="filename" type="text" className={`input ${errors.filename ? 'invalid' : ''}`}
                        value={params.filename} onChange={handleInputChange} disabled={isSimulating} />
                    {errors.filename && <small className="error-text">{errors.filename}</small>}
                </div>

                <div className="input-box">
                    <label htmlFor="size">Nº de Residências</label>
                    <input name="size" type="number" className={`input ${errors.size ? 'invalid' : ''}`}
                        placeholder="Ex: 10000" value={params.size} onChange={handleInputChange} disabled={isSimulating} />
                    {errors.size && <small className="error-text">{errors.size}</small>}
                </div>

                <div className="input-box">
                    <label htmlFor="day">Dias de Simulação</label>
                    <input name="day" type="number" className={`input ${errors.day ? 'invalid' : ''}`}
                        placeholder="Ex: 7" value={params.day} onChange={handleInputChange} disabled={isSimulating} />
                    {errors.day && <small className="error-text">{errors.day}</small>}
                </div>
                
                <div className="input-box">
                    <label htmlFor="toiletType">Tipo de Vaso</label>
                    <select 
                        name="toiletType" 
                        className={`input ${errors.toiletType ? 'invalid' : ''}`}
                        value={params.toiletType} 
                        onChange={handleInputChange} 
                        disabled={isSimulating}
                    >
                        <option value={1}>Vaso 1 (2 Lpf)</option>
                        <option value={2}>Vaso 2 (4.5 Lpf)</option>
                        <option value={3}>Vaso 3 (15 Lpf)</option>
                        <option value={4}>Vaso 4 (3 Lpf)</option>
                    </select>
                    {errors.toiletType && <small className="error-text">{errors.toiletType}</small>}
                </div>

                <div className="input-box">
                    <label htmlFor="showerType">Tipo de Chuveiro</label>
                    <select 
                        name="showerType" 
                        className={`input ${errors.showerType ? 'invalid' : ''}`}
                        value={params.showerType} 
                        onChange={handleInputChange} 
                        disabled={isSimulating}
                    >
                        <option value={1}>Chuveiro 1 (5.66 L/min)(6.17 min)</option>
                        <option value={2}>Chuveiro 2 (5.45 L/min)(6.52 min)</option>
                    </select>
                    {errors.showerType && <small className="error-text">{errors.showerType}</small>}
                </div>

                {/* --- BOTÃO DE AÇÃO --- */}
                <button className="btn" onClick={startSimulation} disabled={isSimulating}>
                    {isSimulating ? 'Simulando...' : 'Iniciar Simulação'}
                </button>
            </div>
            
            {/* --- CAIXA DE STATUS --- */}
            <div className={`result ${isError ? 'error' : ''}`}>
                {statusMessage}
            </div>
        </div>
    );
}

export default SimulationForm;
// --- END OF FILE SimulationForm.jsx ---