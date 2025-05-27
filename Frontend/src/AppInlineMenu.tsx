import React, { useContext, useRef } from 'react';
import { CSSTransition } from 'react-transition-group';
import { classNames } from 'primereact/utils';
import { AuthContext } from './context/auth';

const AppInlineMenu = (props: any) => {
    const menuRef = useRef(null);

    const isSlim = () => {
        return props.menuMode === 'slim';
    };

    const isStatic = () => {
        return props.menuMode === 'static';
    };

    const isSidebar = () => {
        return props.menuMode === 'sidebar';
    };

    const isOverlay = () => {
        return props.menuMode === 'overlay';
    };

    const isMobile = () => {
        return window.innerWidth <= 991;
    };
    
    const { singOut } = useContext(AuthContext);

    return (
        <>
            {!isMobile() && (isStatic() || isSlim() || isSidebar() || isOverlay()) && (
                <div className={classNames('layout-inline-menu', { 'layout-inline-menu-active': props.activeInlineProfile })}>
                    <button className="layout-inline-menu-action p-link" onClick={props.onChangeActiveInlineMenu}>
                        <img  alt="avatar" style={{ width: '44px', height: '44px', objectFit: 'scale-down' }} />
                        <span className="layout-inline-menu-text">Usuário</span>
                        <i className="layout-inline-menu-icon pi pi-angle-down"></i>
                    </button>
                    <CSSTransition nodeRef={menuRef} classNames="p-toggleable-content" timeout={{ enter: 1000, exit: 450 }} in={props.activeInlineProfile} unmountOnExit>
                        <ul ref={menuRef} className="layout-inline-menu-action-panel">
                            <li className="layout-inline-menu-action-item">
                                <button onClick={singOut} className="p-link">
                                    <i className="pi pi-power-off pi-fw"></i>
                                    <span>Sair</span>
                                </button>
                            </li>
                            <li className="layout-inline-menu-action-item">
                                <button className="p-link">
                                    <i className="pi pi-cog pi-fw"></i>
                                    <span>Configurações</span>
                                </button>
                            </li>
                            <li className="layout-inline-menu-action-item">
                                <button className="p-link">
                                    <i className="pi pi-user pi-fw"></i>
                                    <span>Perfil</span>
                                </button>
                            </li>
                        </ul>
                    </CSSTransition>
                </div>
            )}
        </>
    );
};

export default AppInlineMenu;
