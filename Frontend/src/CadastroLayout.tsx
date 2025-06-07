import { Outlet, useNavigate } from "react-router-dom";
import { Button } from "primereact/button";
import './CadastroLayout.css';

export default function CadastroLayout() {
  const navigate = useNavigate();

  return (
    <div className="layout-wrapper">
      <div className="side-menu">
        <img src="/assets/logo.png" alt="Logo" className="logo" />

        <Button label="Registrar Avaliação" icon="pi pi-angle-right" iconPos="right"
          className="menu-button" onClick={() => navigate("/cadastro-avaliacao")} />
        <Button label="Cadastrar novo aluno" icon="pi pi-angle-right" iconPos="right"
          className="menu-button" onClick={() => navigate("/cadastro-aluno")} />
        <Button label="Cadastrar novo Instrutor" icon="pi pi-angle-right" iconPos="right"
          className="menu-button" onClick={() => navigate("/cadastro-instrutor")} />
      </div>

      <div className="layout-content">
        <Outlet />
      </div>
    </div>
  );
}