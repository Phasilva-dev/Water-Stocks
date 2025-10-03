// --- START OF FILE ResultsViewer.jsx ---
import { useState } from 'react';
// Importe as funções do Go
import { SelectFile, GenerateAnalysisAndOpenBrowser } from '../../wailsjs/go/main/App';
import Spinner from '../components/Spinner'; // Importe o novo componente

function ResultsViewer() {
    const [fileName, setFileName] = useState(null);
    const [error, setError] = useState(null);
    const [isProcessing, setIsProcessing] = useState(false);

    const processAndGenerateReport = (filePath) => {
        if (!filePath) return;
        
        setError(null);
        setIsProcessing(true);
        // Extrai o nome do arquivo do caminho completo
        const name = filePath.split(/[\\/]/).pop();
        setFileName(name);

        GenerateAnalysisAndOpenBrowser(filePath)
            .then(successMessage => {
                console.log(successMessage);
                // O estado de processamento é finalizado no `finally` block
            })
            .catch(err => {
                console.error("Erro ao gerar relatório:", err);
                setError(`Falha ao gerar o relatório: ${err}`);
            })
            .finally(() => {
                setIsProcessing(false); // Garante que o estado de processamento termine
            });
    };

    const handleFileSelectClick = () => {
        setError(null);
        // Chama a função do backend para abrir a caixa de diálogo de arquivo
        SelectFile()
            .then(filePath => {
                // Se o usuário selecionar um arquivo, filePath não será vazio
                if (filePath) {
                    processAndGenerateReport(filePath);
                }
            })
            .catch(err => {
                // Wails retorna um erro quando o diálogo é cancelado, vamos ignorá-lo.
                // Mensagens comuns de cancelamento incluem "No file selected" ou "cancelled".
                if (err && !err.toLowerCase().includes('cancel')) {
                    setError(`Erro ao abrir diálogo: ${err}`);
                }
            });
    };
    
    // Gera a mensagem de status a ser exibida
    const renderStatus = () => {
        if (isProcessing) {
            return (
                <div>
                    <p>Analisando dados, por favor aguarde...</p>
                    <Spinner />
                </div>
            );
        }
        if (error) {
            return <p className="drop-zone-error">{error}</p>; // Reutilizando a classe de erro para consistência
        }
        if (fileName) {
            return (
                <p className="drop-zone-success">
                    Pronto para analisar um novo arquivo. Último processado: <strong>{fileName}</strong>
                </p>
            );
        }
        return <p>Selecione um arquivo de análise para começar.</p>;
    };

    return (
        <div className="page-container">
            <h3>Visualizador de Resultados</h3>
            <p>
                Selecione um arquivo de análise (.csv) para gerar um relatório completo em uma nova página do navegador.
            </p>
            
            <button className="btn" onClick={handleFileSelectClick} disabled={isProcessing}>
                {isProcessing ? 'Processando...' : 'Selecionar Arquivo de Análise'}
            </button>
            
            {/* Caixa de status que substitui a antiga drop-zone */}
            <div className={`result ${error ? 'error' : ''}`}>
                {renderStatus()}
            </div>
        </div>
    );
}

export default ResultsViewer;
// --- END OF FILE ResultsViewer.jsx ---