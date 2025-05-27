import { useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import AppMenu from './AppMenu';
import { classNames } from 'primereact/utils';
import { AuthContext } from './context/auth';

const AppTopbar = (props: any) => {
    const navigate = useNavigate();
    const { singOut } = useContext(AuthContext);

    return (
        <>
            <div className="layout-topbar">
                <div className="layout-topbar-left">
                    <button className="logo p-link" onClick={() => navigate('/')}>
                        {/* <img  /> */}
                    </button>

                    <button className="topbar-menu-button p-link" onClick={props.onMenuButtonClick}>
                        <i className="pi pi-bars"></i>
                    </button>
                </div>

                <AppMenu
                    model={props.items}
                    menuMode={props.menuMode}
                    colorScheme={props.colorScheme}
                    menuActive={props.menuActive}
                    activeInlineProfile={props.activeInlineProfile}
                    onSidebarMouseOver={props.onSidebarMouseOver}
                    onSidebarMouseLeave={props.onSidebarMouseLeave}
                    toggleMenu={props.onToggleMenu}
                    onChangeActiveInlineMenu={props.onChangeActiveInlineMenu}
                    onMenuClick={props.onMenuClick}
                    onRootMenuItemClick={props.onRootMenuItemClick}
                    onMenuItemClick={props.onMenuItemClick}
                />

                <div className="layout-topbar-right">
                    <ul className="layout-topbar-right-items">
                        <li id="profile" className={classNames('profile-item', { 'active-topmenuitem': props.topbarMenuActive })}>
                            <button className="p-link" onClick={props.onTopbarItemClick}>
                                <img src="" alt="avatar" style={{ width: '40px', height: '30px' }} />
                            </button>

                            <ul className="fadeInDown">
                                <li role="menuitem">
                                    <button onClick={singOut} className="p-link">
                                        <i className="pi pi-power-off pi-fw"></i>
                                        <span>Sair</span>
                                    </button>
                                </li>
                                <li role="menuitem">
                                    <button className="p-link">
                                        <i className="pi pi-cog pi-fw"></i>
                                        <span>Configurações</span>
                                    </button>
                                </li>
                                <li role="menuitem">
                                    <button className="p-link">
                                        <i className="pi pi-user pi-fw"></i>
                                        <span>Perfil</span>
                                    </button>
                                </li>
                            </ul>
                        </li>
                    </ul>
                </div>
            </div>
        </>
    );
};

export default AppTopbar;
