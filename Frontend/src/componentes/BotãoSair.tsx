import React, { useContext } from 'react';
import { Button } from 'primereact/button';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../context/auth';


interface BotaoSairProps {
  label?: string;
  className?: string;
}

const BotaoSair: React.FC<BotaoSairProps> = ({ label = "Sair", className = "" }) => {
  const navigate = useNavigate();
  const { singOut } = useContext(AuthContext);

  const handleClick = () => {
    singOut();
    navigate('/login', { replace: true });
  };

  return (
    <Button 
      label={label} 
      icon="pi pi-sign-out" 
      className={`p-button-secondary ${className}`} 
      onClick={handleClick} 
    />
  );
};

export default BotaoSair;
