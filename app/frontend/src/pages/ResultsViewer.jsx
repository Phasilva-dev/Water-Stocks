// --- START OF FILE ResultsViewer.jsx ---
import { useState } from 'react';
// Importe as funções do Go
import { SelectFile, GenerateAnalysisAndOpenBrowser } from '../../wailsjs/go/main/App';
import Spinner from '../components/Spinner'; // Importe o novo componente

function ResultsViewer() {
    const [fileName, setFileName] = useState(null);
    const [error, setError] = useState(null);
    const [isProcessing, setIsProcessing] = useState(false);
    const [isDragOver, setIsDragOver] = useState(false);

    const processAndGenerateReport = (filePath) => {
        if (!filePath) return;
        
        setError(null);
        setIsProcessing(true);
        const name = filePath.split(/[\\/]/).pop();
        setFileName(name);

        GenerateAnalysisAndOpenBrowser(filePath)
            .then(successMessage => {
                console.log(successMessage);
                // A mensagem de sucesso ou erro é tratada no Go, aqui apenas resetamos o estado
                setIsProcessing(false);
            })
            .catch(err => {
                console.error("Erro ao gerar relatório:", err);
                setError(`Falha ao gerar o relatório: ${err}`);
                setIsProcessing(false);
            });
    };

    const handleFileSelectClick = () => {
        setError(null);
        SelectFile()
            .then(filePath => {
                // Se o usuário cancelar, filePath será uma string vazia
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

    const handleDragOver = (e) => { e.preventDefault(); e.stopPropagation(); setIsDragOver(true); };
    const handleDragLeave = (e) => { e.preventDefault(); e.stopPropagation(); setIsDragOver(false); };
    const handleDrop = (e) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragOver(false);
        const file = e.dataTransfer.files?.[0];
        
        if (!file || !file.name.toLowerCase().endsWith('.csv')) {
            setError('Erro: Por favor, selecione apenas arquivos .csv.');
            setFileName(null);
            return;
        }
        // Wails usa `file.path` para acessar o caminho completo do arquivo.
        processAndGenerateReport(file.path);
    };

    return (
        <div className="page-container">
            <h3>Visualizador de Resultados</h3>
            <p>Selecione um arquivo de análise (.csv) ou arraste e solte na área abaixo para gerar um relatório completo em uma nova página do navegador.</p>
            
            <button className="btn" onClick={handleFileSelectClick} disabled={isProcessing}>
                {isProcessing ? 'Processando...' : 'Selecionar Arquivo de Análise'}
            </button>
            
            <div 
                className={`drop-zone ${isDragOver ? 'drag-over' : ''}`}
                onDragOver={handleDragOver}
                onDragLeave={handleDragLeave}
                onDrop={handleDrop}
            >
                {isProcessing ? (
                    <div>
                        <p>Analisando dados, por favor aguarde...</p>
                        <Spinner />
                    </div>
                ) : error ? (
                    <p className="drop-zone-error">{error}</p>
                ) : fileName ? (
                    <p className="drop-zone-success">Pronto para analisar um novo arquivo. Último processado: <strong>{fileName}</strong></p>
                ) : (
                    <p>Arraste e solte um arquivo .csv aqui</p>
                )}
            </div>
        </div>
    );
}

export default ResultsViewer;
// --- END OF FILE ResultsViewer.jsx ---