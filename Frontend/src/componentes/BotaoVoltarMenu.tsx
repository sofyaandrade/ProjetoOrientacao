import React from 'react';
import { Button } from 'primereact/button';
import { useNavigate } from 'react-router-dom';

interface BotaoVoltarMenuProps {
  label?: string;
  className?: string;
}

const BotaoVoltarMenu: React.FC<BotaoVoltarMenuProps> = ({ label = "Voltar para o Menu", className = "" }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate('/menu-principal');
  };

  return (
    <Button 
      label={label} 
      icon="pi pi-arrow-left" 
      className={`p-button-secondary ${className}`} 
      onClick={handleClick} 
    />
  );
};

export default BotaoVoltarMenu;