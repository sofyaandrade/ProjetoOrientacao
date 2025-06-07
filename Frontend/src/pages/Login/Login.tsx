import React, { useState, useRef } from 'react';
import { InputText } from 'primereact/inputtext';
import { Button } from 'primereact/button';
import { Navigate } from 'react-router-dom';
import useAuth from '../../hooks/useAuth';
import './Login.css';
import { Toast } from 'primereact/toast';

export const Login = (props: any) => {

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const { signIn, signed } = useAuth();
    const toast = useRef<any>(null);
    const ERRO_BUSCA_USUARIO = 404
	const ERRO_SENHA_INCORRETA = 401

    const handleSubmit = async (e: any) => {
        e.preventDefault();
        const data = {
            email,
            password,
        };
        const ResponseSignIn = await signIn(data);

        if(ResponseSignIn === ERRO_BUSCA_USUARIO){
            toast.current.show({severity: 'error', summary: 'Erro', detail: `Email inválido: ${email}`, life: 3000,
              });
        } else if (ResponseSignIn === ERRO_SENHA_INCORRETA){
            toast.current.show({ severity: 'error', summary: 'error', detail: 'Senha incorreta', life: 3000 });
        }
    };

    const handleKeyUp = (e: any) => {
        if(e.key === 'Enter') {
            handleSubmit(e)
        }
    };


    if (signed) {
        return <Navigate to="/menu-principal"></Navigate>
    } else {

        return (
            <div className="login-body">
                <Toast ref={toast} />
                <div className="login-panel p-fluid ajuste-login">
                    <div className="flex flex-column align-items-center">
                        <div className="flex align-items-center mb-2 logo-container">
                        </div>
                        <div className="form-container">
                            <span className="p-input-icon-left">
                                <i className="pi pi-envelope"></i>
                                <InputText value={email} type="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)} />
                            </span>
                            <span className="p-input-icon-left">
                                <i className="pi pi-key"></i>
                                <InputText value={password} onChange={(e) => setPassword(e.target.value)} type="password" placeholder="Senha" onKeyUp={handleKeyUp}/>
                            </span>
                            <button className="flex p-link">Esqueceu senha?</button>
                        </div>
                        <div className="button-container">
                            <Button type="button" className="botao-login" onClick={handleSubmit} label={'Entrar'}></Button>
                        </div>
                    </div>

                    <div className="login-footer flex align-items-center">
                        <div className="footer-logo-container">
                           
                        </div>
                    </div>
                </div>
            </div>
        );
    }
};
