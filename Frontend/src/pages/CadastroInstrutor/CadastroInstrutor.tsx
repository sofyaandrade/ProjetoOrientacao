import { InputText } from "primereact/inputtext"
import { Toast } from "primereact/toast"
import { Accordion, AccordionTab } from "primereact/accordion";
import { Button } from "primereact/button";
import { useEffect, useRef, useState } from "react";
import { useAppDispatch, useAppSelector } from "../../store/hooks";
import IAvaliacao from "../../interface/IAvaliacao/IAvaliacao";
import { addAvaliacao, getAvaliacao, updateAvaliacao } from "../../service/avaliacaoService";
import IUsuario from "../../interface/IUsuario/IUsuario";
import { Toolbar } from "primereact/toolbar";
import { DataTable } from "primereact/datatable";
import { Column } from "primereact/column";
import { Dialog } from "primereact/dialog";
import { addUsuario, deleteUsuario, getUsuarios, updateUsuario } from "../../service/usuarioService";
import { getUsuarioPerfil } from "../../service/usuarioPerfilService";
import { fetchUsuarioLogado } from "../../store/reducers/usuarioLogado";
import IPerfil from "../../interface/IUsuario/IPerfil";
import { Dropdown } from "primereact/dropdown";
import { Checkbox } from "primereact/checkbox";
import { InputMask } from "primereact/inputmask";
import { ToggleButton } from "primereact/togglebutton";
import { classNames } from 'primereact/utils';
import BotaoVoltarMenu from "../../componentes/BotaoVoltarMenu";
import BotaoSair from "../../componentes/BotãoSair";

export function CadastroUsuario() {
    const labelUsuario = 'Usuário'
    const USUARIO_ADMINITRADOR_PERFIL_ID = 1
    
    const listaUsuariosStore = useAppSelector(state => state.usuario)
    const dispatch = useAppDispatch()

    let usuarioVazio: IUsuario
    usuarioVazio = {
        ID: 0,
        Nome: "",
        Telefone: 0,
        Email: "",
        Password: "",
        UsuarioPerfilID: 0,
        UsuarioPerfil: {
            ID: 0,
            Descricao: "",
            Usuarios: [],
        },
    }


    const [listaUsuarios, setListaUsuarios] = useState<IUsuario[]>();
    const [listaUsuarioPerfil, setListaUsuarioPerfil] = useState<IPerfil[]>();
    const [usuario, setUsuario] = useState<IUsuario>(usuarioVazio);
    const [usuarioDialog, setUsuarioDialog] = useState<boolean>(false);
    const [usuarioLogado, setUsuarioLogado] = useState<IUsuario>()
    const [editarUsuario, setEditarUsuario] = useState<boolean>(false);
    const [deleteUsuarioDialog, setDeleteUsuarioDialog] = useState<boolean>(false);
    const [submitted, setSubmitted] = useState<boolean>(false);
    const [checked, setChecked] = useState<boolean>(false);
    const [globalFilter, setGlobalFilter] = useState<any>(null);
    const toast = useRef<any>(null);
    const [atualizaSenhaCheckbox, setAtualizaSenhaCheckbox] = useState(false);

    useEffect(() => {
        dispatch(getUsuarios()).then((data) => setListaUsuarios(data.payload));
        dispatch(getUsuarioPerfil()).then((data) => setListaUsuarioPerfil(data.payload));
        dispatch(fetchUsuarioLogado()).then((data) => setUsuarioLogado(data.payload))
    }, [])

    useEffect(() => {
        dispatch(getUsuarios()).then((data) => setListaUsuarios(data.payload))
    }, [listaUsuarios])

    const openNew = () => {
        setAtualizaSenhaCheckbox(true);
        setUsuario(usuarioVazio);
        setSubmitted(false);
        setUsuarioDialog(true);
    };

    const hideDialog = () => {
        setSubmitted(false);
        setUsuarioDialog(false);
        setEditarUsuario(false);
    };

    const hideDeleteUsuarioDialog = () => {
        setDeleteUsuarioDialog(false);
    };

    const verificaEmailExistente = (email: string, listaUsuarios: IUsuario[] | undefined, usuarioAtualID?: number) => {
        if (!listaUsuarios) return false; 
        return listaUsuarios.some(u => u.Email === email && u.ID !== usuarioAtualID);
    };

    const saveUsuario = () => {
        setSubmitted(true);
        if (usuario.UsuarioPerfilID !== USUARIO_ADMINITRADOR_PERFIL_ID || usuarioLogado?.UsuarioPerfilID === USUARIO_ADMINITRADOR_PERFIL_ID) {
            if (usuario.Nome.trim() && usuario.Email.trim() && (atualizaSenhaCheckbox ? usuario.Password.trim() : true) && usuario.Telefone !== 0) {                
                let _usuario: any = { ...usuario };
                if (verificaEmailExistente(usuario.Email, listaUsuarios, usuario.ID)) {
                    toast.current.show({ severity: 'warn', summary: 'Aviso' , detail: 'Email já cadastrado' , life: 3000});
                    return;
                }
                    if (editarUsuario) {    
                        dispatch(updateUsuario(_usuario))    
                        toast.current.show({ severity: 'success', summary: 'Sucesso', detail: 'Atualizado com sucesso', life: 3000 });
                    } else {
                        dispatch(addUsuario(usuario))
                        toast.current.show({ severity: 'success', summary: 'Sucesso', detail: 'Cadastrado com sucesso', life: 3000 });
                    }
                    dispatch(getUsuarios()).then((data) => setListaUsuarios(data.payload));
                    setEditarUsuario(false);
                    setUsuarioDialog(false);
                    setUsuario(usuarioVazio);
            }
        } else {
            toast.current.show({severity: 'warn', summary: 'Aviso', detail: 'Não cadastrar usuarios Administrador' , life: 3000 });
        }
    };

    const editUsuario = (usuario: any) => {
        setUsuario({ ...usuario });
        setEditarUsuario(true);
        setUsuarioDialog(true);
        setChecked(usuario.Telefone.toString().length === 10)
        
    };

    const confirmDeleteUsuario = (usuario: IUsuario) => {
        if (usuario.UsuarioPerfilID !== USUARIO_ADMINITRADOR_PERFIL_ID || usuarioLogado?.UsuarioPerfilID === USUARIO_ADMINITRADOR_PERFIL_ID) {
            setUsuario(usuario);
            setDeleteUsuarioDialog(true);
        } else {
            toast.current.show({severity: 'warn', summary: 'Aviso' , detail: 'Não exclui usuário adminitrador', life: 3000 });
        }       
    };

    const excluirUsuario = () => {

        dispatch(deleteUsuario(usuario.ID))
        dispatch(getUsuarios()).then((data) => setListaUsuarios(data.payload));

        setDeleteUsuarioDialog(false);
        setUsuario({ ...usuarioVazio });
        toast.current.show({ severity: 'success', summary: 'Sucesso', detail: 'Excluído com sucesso' , life: 3000 });
    };

    const onTelefoneChange = (e: any) => {
        let _usuario = { ...usuario };
        const telefone = e.target.value.replace(/\D/g, ''); 
        _usuario.Telefone = Number(telefone);
        setUsuario(_usuario);
    };

    const onPerfilChange = (e: any) => {
        let _usuario = { ...usuario };
        _usuario.UsuarioPerfil = e.value;
        _usuario.UsuarioPerfilID = _usuario.UsuarioPerfil.ID
        setUsuario({ ..._usuario });
    };



    const leftToolbarTemplate = () => {
        return (
            <>
                <Button label={'Novo'} icon="pi pi-plus" className="p-button-success mr-2 mb-2" onClick={openNew} />
            </>
        );
    };

    const onTelefoneFocus = (e: any) => {
        if (usuario.Telefone === 0) {
            let _usuario = { ...usuario };
            _usuario.Telefone = 0;
            setUsuario(_usuario);
        }
        e.target.setSelectionRange(0, 0);
    };
    

    const idBodyTemplate = (rowData: IUsuario) => {
        return rowData.ID;
    };

    const nomeBodyTemplate = (rowData: IUsuario) => {
        return rowData.Nome;
    };

   function formatarTelefone(str: string) {
        if (str.length === 10) {
            let match = str.match(/^(\d{2})(\d{4})(\d{4})$/);
            if (match) {
                return '(' + match[1] + ') ' + match[2] + '-' + match[3];
            }
        } else {
            let match = str.match(/^(\d{2})(\d{5})(\d{4})$/);
            if (match) {
                return '(' + match[1] + ') ' + match[2] + '-' + match[3];
            }
        }
        return '';
    }

      const telefoneBodyTemplate = (rowData: IUsuario) => {
            let telefone = formatarTelefone(rowData?.Telefone.toString())
            return (
                <>
                    {telefone}
                </>
            );
        };

    const emailBodyTemplate = (rowData: IUsuario) => {
        return rowData.Email;
    };

    const roleBodyTemplate = (rowData: IUsuario) => {
        return rowData.UsuarioPerfil.Descricao;
    };

    const actionBodyTemplate = (rowData: any) => {
        return (
            <div className="flex align-items-end flex-wrap justify-content-center actions gap-2">
                <Button icon="pi pi-pencil" className="p-button-rounded p-button-success mr-2" onClick={() => editUsuario(rowData)} />
                <Button icon="pi pi-trash" className="p-button-rounded p-button-warning mt-2" onClick={() => confirmDeleteUsuario(rowData)} />
            </div>
        );
    };


    const header = (
        <div className="flex flex-column md:flex-row md:justify-content-between md:align-items-center">
            <h5 className="m-0">{'Listagem de Usuários'}</h5>
        </div>
    );

    const usuarioDialogFooter = (
        <>
            <Button label={'Cancelar'} icon="pi pi-times" className="p-button-text" onClick={hideDialog} />
            <Button label={'Salvar'} icon="pi pi-check" className="p-button-text" onClick={saveUsuario} />
        </>
    );
    const deleteUsuarioDialogFooter = (
        <>
            <Button label={'Não'} icon="pi pi-times" className="p-button-text" onClick={hideDeleteUsuarioDialog} />
            <Button label={'Sim'} icon="pi pi-check" className="p-button-text" onClick={excluirUsuario} />
        </>
    );

    const telefoneMask = checked ? "(99) 9999-9999" : "(99) 99999-9999";
    
    return (
        <>
            <h2>{'Cadastro de Usuarios'}</h2>

            <div className="grid crud-demo">
                <div className="col-12">
                    <div className="card">
                        <Toast ref={toast} />
                        <Toolbar className="mb-4" left={leftToolbarTemplate}></Toolbar>

                        <DataTable
                            value={listaUsuarios}
                            dataKey="id"
                            paginator
                            rows={10}
                            rowsPerPageOptions={[5, 10, 25]}
                            className="datatable-responsive"
                            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown CurrentPageReport"
                            currentPageReportTemplate={'{first} - {last} ' + ' de ' + ' {totalRecords}'}
                            globalFilter={globalFilter}
                            emptyMessage={'Nenhum Instrutor Cadastrado'}
                            header={header}
                        >
                            <Column field="ID" header={'Id'} sortable body={idBodyTemplate} headerStyle={{ width: '10%', minWidth: '4rem' }}></Column>
                            <Column field="nome" header={'Nome'} sortable body={nomeBodyTemplate} headerStyle={{ width: '30%', minWidth: '8rem' }}></Column>
                            <Column field="telefone" header={'Telefone'} sortable   body={telefoneBodyTemplate} headerStyle={{ width: '20%', minWidth: '8rem' }}></Column>
                            <Column field="email" header={'Email'} sortable body={emailBodyTemplate} headerStyle={{ width: '20%', minWidth: '10rem' }}></Column>
                            <Column field="UsuarioPerfilID" header={'Role'} sortable body={roleBodyTemplate} headerStyle={{ width: '20%', minWidth: '10rem' }}></Column>
                            <Column headerStyle={{ width: '10%', minWidth: '10%' }} body={actionBodyTemplate}></Column>
                        </DataTable>

                        <Dialog visible={usuarioDialog} style={{ width: '450px' }} header={'Detalhes usuario'} modal className="p-fluid" footer={usuarioDialogFooter} onHide={hideDialog}>
                            <div className="field">
                                <label htmlFor="nome">{'Nome'}</label>
                                <InputText id="nome" value={usuario.Nome} required autoFocus onChange={(e) => setUsuario({...usuario, Nome: e.target.value})} className={classNames({ 'p-invalid': submitted && !usuario.Nome })} />
                                {submitted && !usuario.Nome && <small className="p-invalid">{'Nome obrigatório'}</small>}
                            </div>
                            <div className="formgrid grid">
                                <div className="field col">
                                    <label htmlFor="telefone">{'telefone'}</label>
                                    <div className="mb-1 mt-1">
                                        <ToggleButton onLabel="Fixo" offLabel="Celular" onIcon="pi pi-phone" offIcon="pi pi-mobile" checked={checked} onChange={(e) => {setChecked(e.value); setUsuario({...usuario, Telefone: 0}); }} />
                                    </div>
                                    <InputMask id="telefone" mask={telefoneMask} value={usuario.Telefone === 0 ? "" : usuario.Telefone.toString()} onChange={onTelefoneChange} onFocus={onTelefoneFocus} placeholder={checked ? "(99) 9999-9999" : "(99) 99999-9999"} unmask={true} />
                                        {submitted && !usuario.Telefone && (
                                        <small className="p-invalid">{'Telefone obrigatório'}</small>
                                    )}
                                </div>
                            </div>
                            <div className="field">
                                <label htmlFor="email">{'Email'}</label>
                                <InputText id="email" value={usuario.Email} required onChange={(e) => setUsuario({...usuario, Email: e.target.value})} className={classNames({ 'p-invalid': submitted && !usuario.Email })} />
                                {submitted && !usuario.Email && <small className="p-invalid">{'Email obrigatório'}</small>}
                            </div>

                            <div className="field">
                                <div className='mb-2'>
                                    {editarUsuario && (
                                        <Checkbox inputId="senha" onChange={() => setAtualizaSenhaCheckbox((prevChecked) => !prevChecked)} checked={atualizaSenhaCheckbox} className={'mr-2'}></Checkbox>
                                    )}

                                    <label htmlFor="senha">
                                        {editarUsuario ? 'Atualizar senha' : 'Senha'}
                                    </label>
                                </div>

                                {atualizaSenhaCheckbox && (
                                    <>
                                        <InputText id="password" value={usuario.Password} required onChange={(e) => setUsuario({ ...usuario, Password: e.target.value })} className={classNames({ 'p-invalid': submitted && !usuario.Password })}/>
                                        {submitted && !usuario.Password && (
                                            <small className="p-invalid">{'Senha obrigatório'}</small>
                                        )}
                                    </>
                                )}
                            </div>

                            <div className="field">
                                <label className="mb-3">{'Perfil'}</label>
                                <div className="formgrid grid">
                                    <Dropdown value={usuario.UsuarioPerfil} options={listaUsuarioPerfil} optionLabel="Descricao" onChange={(e) => onPerfilChange(e)} placeholder="Selecione o tipo" />
                                </div>
                            </div>

                        </Dialog>

                        <Dialog visible={deleteUsuarioDialog} style={{ width: '450px' }} header={'Confirmar'} modal footer={deleteUsuarioDialogFooter} onHide={hideDeleteUsuarioDialog}>
                            <div className="flex align-items-center justify-content-center">
                                <i className="pi pi-exclamation-triangle mr-3" style={{ fontSize: '2rem' }} />
                                {usuario && (
                                    <span>
                                        <b>Deseja excluir?</b>
                                    </span>
                                )}
                            </div>
                        </Dialog>

                        <BotaoVoltarMenu />
                        <BotaoSair />

                    </div>
                </div>
            </div>

        </>
    );
}