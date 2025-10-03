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
        if (!params.filename) newErrors.filename = "Nome do arquivo é obrigatório.";
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
            return;
        }

        setIsSimulating(true);
        setStatusMessage(`Iniciando simulação '${params.filename}'...`);
        
        RunSimulation(params.size, params.day, params.toiletType, params.showerType, params.filename)
            .then(response => setStatusMessage(response))
            .catch(err => setStatusMessage(`Erro: ${err}`));
    };

    // --- EFEITO PARA OUVIR EVENTOS DO BACKEND (WAILS) ---
    useEffect(() => {
        const cleanup = window.runtime.EventsOn('simulationComplete', (message) => {
            setStatusMessage(message);
            setIsSimulating(false); // Libera o botão
        });
        return cleanup;
    }, []);

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
                        <option value={1}>Vaso Padrão (6 Lpf)</option>
                        <option value={2}>Vaso com Caixa Acoplada (3/6 Lpf)</option>
                        <option value={3}>Vaso a Vácuo (1.2 Lpf)</option>
                        <option value={4}>Vaso Ecológico (4.8 Lpf)</option>
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
                        <option value={1}>Chuveiro Comum (12 L/min)</option>
                        <option value={2}>Chuveiro Econômico (8 L/min)</option>
                    </select>
                    {errors.showerType && <small className="error-text">{errors.showerType}</small>}
                </div>

                {/* --- BOTÃO DE AÇÃO --- */}
                <button className="btn" onClick={startSimulation} disabled={isSimulating}>
                    {isSimulating ? 'Simulando...' : 'Iniciar Simulação'}
                </button>
            </div>
            
            {/* --- CAIXA DE STATUS --- */}
            <div className={`result ${Object.keys(errors).length > 0 ? 'error' : ''}`}>
                {statusMessage}
            </div>
        </div>
    );
}

export default SimulationForm;
// --- END OF FILE SimulationForm.jsx ---