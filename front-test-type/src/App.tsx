import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Navigate, useNavigate } from 'react-router-dom';
import { Button } from 'primereact/button';

function App() {
  const navigate = useNavigate();

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <div className="main-container">
      <img src="/assets/logo.png" alt="Logo" className="logo" />

      <div className="button-group">
        <Button
          label="Registrar Avaliação"
          icon="pi pi-angle-right"
          iconPos="right"
          className="menu-button"
          onClick={() => <Navigate to="/avaliacao" replace />}
        />
        <Button
          label="Cadastrar novo aluno"
          icon="pi pi-angle-right"
          iconPos="right"
          className="menu-button"
          onClick={() => navigate("/novo-aluno")}
        />
        <Button
          label="Cadastrar novo Instrutor"
          icon="pi pi-angle-right"
          iconPos="right"
          className="menu-button"
          onClick={() => navigate("/novo-instrutor")}
        />
      </div>
    </div>
        <Navigate to="/login" replace />;
        
      </header>
    </div>
  );
}

export default App;
