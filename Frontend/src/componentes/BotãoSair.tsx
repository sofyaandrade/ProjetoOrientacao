import React from 'react';
import { Button } from 'primereact/button';
import { useNavigate } from 'react-router-dom';

interface BotaoVoltarMenuProps {
  label?: string;
  className?: string;
}

const BotaoSair: React.FC<BotaoVoltarMenuProps> = ({ label = "Sair", className = "" }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate('/login');
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

export default BotaoSair;