// src/components/MainMenu.tsx
import { useNavigate } from "react-router-dom";
import { Button } from "primereact/button";
import "./MainMenu.css";
import BotaoSair from "./componentes/BotãoSair";

export default function MainMenu() {
  const navigate = useNavigate();

  return (
    <div className="main-container">
      <img src="/assets/logo.png" alt="Logo" className="logo" />

      <div className="button-group">
        <Button
          label="Registrar Avaliação"
          icon="pi pi-angle-right"
          iconPos="right"
          className="menu-button"
          onClick={() => navigate("/cadastro-avaliacao")}
        />
        <Button
          label="Cadastrar novo aluno"
          icon="pi pi-angle-right"
          iconPos="right"
          className="menu-button"
          onClick={() => navigate("/cadastro-aluno")}
        />
        <Button
          label="Cadastrar novo Instrutor"
          icon="pi pi-angle-right"
          iconPos="right"
          className="menu-button"
          onClick={() => navigate("/cadastro-instrutor")}
        />
        <BotaoSair/>
      </div>
    </div>
  );
}
