const isStandalone = process.env.REACT_APP_MODE === 'STANDALONE';
const apiPort = process.env.REACT_APP_API_PORT;
const apiHost = process.env.REACT_APP_API_HOST;
const apiRelatorioPort = process.env.REACT_APP_RELATORIO_API_PORT;
const apiRelatorioHost = process.env.REACT_APP_RELATORIO_API_HOST;
const gerador = process.env.REACT_APP_GERADOR;

export { isStandalone, apiPort, apiHost, apiRelatorioPort, apiRelatorioHost, gerador };