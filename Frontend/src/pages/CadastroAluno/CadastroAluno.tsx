import { useAppSelector, useAppDispatch } from '../../store/hooks'
import { useEffect, useRef, useState } from 'react';
import IAluno from '../../interface/IAluno/IAluno';
import { classNames } from 'primereact/utils';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { Toast } from 'primereact/toast';
import { Button } from 'primereact/button';
import { Toolbar } from 'primereact/toolbar';
import { InputMask } from 'primereact/inputmask';
import { Dialog } from 'primereact/dialog';
import { InputText } from 'primereact/inputtext';
import { ToggleButton } from 'primereact/togglebutton';
import { Dropdown } from 'primereact/dropdown';
import { Checkbox } from 'primereact/checkbox';
import IAvaliacao from '../../interface/IAvaliacao/IAvaliacao';
import ICliente from '../../interface/IAluno/IAluno';
import { addAlunos, deleteAlunos, getAlunos, updateAlunos } from '../../service/alunoService';
import BotaoVoltarMenu from '../../componentes/BotaoVoltarMenu';


export function CadastroAluno() {
    const USUARIO_ADMIN_ID = 1
    const dispatch = useAppDispatch()
   
    let alunoVazio: IAluno
    alunoVazio = {
        ID: 0,
        Nome: "",
        Email: "",
        Telefone: 0,
        Descricao: "",
        Avaliacao: {
            ID: 0,
            Abdomen: 0,
            AbdomenMasc: 0,
            AnteBracoDireito: 0,
            AnteBracoEsquerdo: 0,
            BracoContraidoDireito: 0,
            BracoContraidoEsquerdo: 0,
            BracoRelaxadoDireito: 0,
            BracoRelaxadoEsquerdo: 0,
            Cintura: 0,
            Coxa: 0,
            CoxaDireita: 0,
            CoxaEsquerda: 0,
            PernaDireita: 0,
            PernaEsquerda: 0,
            Quadril: 0,
            SupraIliaca: 0,
            Torax: 0,
            ToraxExpirado: 0,
            ToraxInspirado: 0,
            ToraxNormal: 0,
            Triceps: 0,
            AlunoID: 0,
        }
    }


    const [submitted, setSubmitted] = useState<boolean>(false)
    const [listaAlunos, setListaAlunos] = useState<IAluno[]>();
    const [aluno, setAluno] = useState<ICliente>(alunoVazio);
    const [alunoDialog, setAlunoDialog] = useState<boolean>(false);
    const [alunoLogado, setAlunoLogado] = useState<IAluno>()
    const [editarAluno, setEditarAluno] = useState<boolean>(false);
    const [deleteAlunoDialog, setDeleteAlunoDialog] = useState<boolean>(false);
    const [checked, setChecked] = useState<boolean>(false);
    const [globalFilter, setGlobalFilter] = useState<any>(null);
    const toast = useRef<any>(null);
    const [atualizaSenhaCheckbox, setAtualizaSenhaCheckbox] = useState(false);
       
//não ta trazendo as infos na primeira carregada
    useEffect(() => {
        dispatch(getAlunos()).then((data) => setListaAlunos(data.payload));
    }, [setListaAlunos])

    const openNew = () => {
        setAtualizaSenhaCheckbox(true);
        setAluno(alunoVazio);
        setSubmitted(false);
        setAlunoDialog(true);
    };

    const hideDialog = () => {
        setSubmitted(false);
        setAlunoDialog(false);
        setEditarAluno(false);
    };

    const hideDeleteAlunoDialog = () => {
        setDeleteAlunoDialog(false);
    };

    const verificaEmailExistente = (email: string, listaAlunos: IAluno[] | undefined) => {
        if (!listaAlunos) return false; 
        return listaAlunos.some(u => u.Email === email);
    };

    const saveAluno = () => {
        setSubmitted(true);
        
        if (aluno.Nome.trim() && aluno.Email.trim() && aluno.Telefone !== 0) {                
            let _aluno: any = { ...aluno };
            if (verificaEmailExistente(aluno.Email, listaAlunos)) {
                toast.current.show({ severity: 'warn', summary: 'Aviso', detail: 'Email já cadastrado', life: 3000});
                return;
            }
                if (editarAluno) {    
                    dispatch(updateAlunos(_aluno))    
                    toast.current.show({ severity: 'success', summary: 'Sucesso', detail: 'Aluno atualizado com sucesso', life: 3000 });
                } else {
                    dispatch(addAlunos(aluno))
                    toast.current.show({ severity: 'success', summary: 'Sucesso', detail: 'Aluno cadastrado com sucesso', life: 3000 });
                }
                dispatch(getAlunos()).then((data) => setListaAlunos(data.payload));
                setEditarAluno(false);
                setAlunoDialog(false);
                setAluno(alunoVazio);
        }
    };

    const editAluno = (aluno: any) => {
        setAluno({ ...aluno });
        setEditarAluno(true);
        setAlunoDialog(true);
        setChecked(aluno.Telefone.toString().length === 10)
    };

    const confirmDeleteAluno = (aluno: IAluno) => {
        setAluno(aluno);
        setDeleteAlunoDialog(true);   
    };

    const excluirAluno = () => {

        dispatch(deleteAlunos(aluno.ID))
        dispatch(getAlunos()).then((data) => setListaAlunos(data.payload));

        setDeleteAlunoDialog(false);
        setAluno({ ...alunoVazio });
        toast.current.show({ severity: 'success', summary: 'Sucesso', detail: 'Aluno excluído com sucesso', life: 3000 });
    };

    const onTelefoneChange = (e: any) => {
        let _aluno = { ...aluno };
        const telefone = e.target.value.replace(/\D/g, ''); 
        _aluno.Telefone = Number(telefone);
        setAluno(_aluno);
    };

    // const onPerfilChange = (e: any) => {
    //     let _aluno = { ...aluno };
    //     _aluno.AlunoPerfil = e.value;
    //     _aluno.AlunoPerfilID = _aluno.AlunoPerfil.ID
    //     setAluno({ ..._aluno });
    // };



    const leftToolbarTemplate = () => {
        return (
            <>
                <Button label={"Novo"} icon="pi pi-plus" className="p-button-success mr-2 mb-2" onClick={openNew} />
            </>
        );
    };

    const onTelefoneFocus = (e: any) => {
        if (aluno.Telefone === 0) {
            let _aluno = { ...aluno };
            _aluno.Telefone = 0;
            setAluno(_aluno);
        }
        e.target.setSelectionRange(0, 0);
    };
    

    const idBodyTemplate = (rowData: IAluno) => {
        return rowData.ID;
    };

    const NomeBodyTemplate = (rowData: IAluno) => {
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

    const telefoneBodyTemplate = (rowData: IAluno) => {
        let telefone = formatarTelefone(rowData?.Telefone.toString())
        return (
            <>
                {telefone}
            </>
        );
    };

    const emailBodyTemplate = (rowData: IAluno) => {
        console.log(rowData.Email)
        return rowData.Email;
    };

    const roleBodyTemplate = (rowData: IAluno) => {
        return rowData.Descricao;
    };

    const actionBodyTemplate = (rowData: any) => {
        return (
            <div className="flex align-items-end flex-wrap justify-content-center actions gap-2">
                <Button icon="pi pi-pencil" className="p-button-rounded p-button-success mr-2" onClick={() => editAluno(rowData)} />
                <Button icon="pi pi-trash" className="p-button-rounded p-button-warning mt-2" onClick={() => confirmDeleteAluno(rowData)} />
            </div>
        );
    };


    const header = (
        <div className="flex flex-column md:flex-row md:justify-content-between md:align-items-center">
            <h5 className="m-0">Listagem Alunos</h5>
        </div>
    );

    const alunoDialogFooter = (
        <>
            <Button label={'Cancelar'} icon="pi pi-times" className="p-button-text" onClick={hideDialog} />
            <Button label={'Salvar'} icon="pi pi-check" className="p-button-text" onClick={saveAluno} />
        </>
    );
    const deleteAlunoDialogFooter = (
        <>
            <Button label={'Não'} icon="pi pi-times" className="p-button-text" onClick={hideDeleteAlunoDialog} />
            <Button label={'Sim'} icon="pi pi-check" className="p-button-text" onClick={excluirAluno} />
        </>
    );

    const telefoneMask = checked ? "(99) 9999-9999" : "(99) 99999-9999";
    

    return (
        <>
            <h2>Cadastro de Alunos</h2>

            <div className="grid crud-demo">
                <div className="col-12">
                    <div className="card">
                        <Toast ref={toast} />
                        <Toolbar className="mb-4" left={leftToolbarTemplate}></Toolbar>

                        <DataTable
                            value={listaAlunos}
                            dataKey="id"
                            paginator
                            rows={10}
                            rowsPerPageOptions={[5, 10, 25]}
                            className="datatable-responsive"
                            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown CurrentPageReport"
                            currentPageReportTemplate={'{first} - {last} ' + 'de' + ' {totalRecords}'}
                            globalFilter={globalFilter}
                            emptyMessage={'Nenhum aluno cadastrado'}
                            header={header}
                        >
                            <Column field="ID" header={'ID'} sortable  headerStyle={{ width: '3%', minWidth: '4rem' }}></Column>
                            <Column field="Nome" header={'Nome'} sortable headerStyle={{ width: '15%', minWidth: '8rem' }}></Column>
                            <Column field="Telefone" header={'Telefone'} sortable body={telefoneBodyTemplate} headerStyle={{ width: '5%', minWidth: '8rem' }}></Column>
                            <Column field="Email" header={'Email'} sortable body={emailBodyTemplate} headerStyle={{ width: '10%', minWidth: '10rem' }}></Column>
                            <Column field="Descricao" header={'Descricao'} sortable body={roleBodyTemplate} headerStyle={{ width: '20%', minWidth: '10rem' }}></Column>
                            <Column headerStyle={{ width: '10%', minWidth: '10%' }} body={actionBodyTemplate}></Column>
                        </DataTable>

                        <Dialog visible={alunoDialog} style={{ width: '450px' }} header={'Detalhes Aluno'} modal className="p-fluid" footer={alunoDialogFooter} onHide={hideDialog}>
                            <div className="field">
                                <label htmlFor="Nome">Nome</label>
                                <InputText id="Nome" value={aluno.Nome} required autoFocus onChange={(e) => setAluno({...aluno, Nome: e.target.value})} className={classNames({ 'p-invalid': submitted && !aluno.Nome })} />
                                {submitted && !aluno.Nome && <small className="p-invalid">{'Nome obrigatório'}</small>}
                            </div>
                            <div className="formgrid grid">
                                <div className="field col">
                                    <label htmlFor="telefone">Telefone</label>
                                    <div className="mb-1 mt-1">
                                    </div>
                                        <InputMask id="telefone" mask={telefoneMask} value={aluno.Telefone === 0 ? "" : aluno.Telefone.toString()} onChange={onTelefoneChange} onFocus={onTelefoneFocus} placeholder={checked ? "(99) 9999-9999" : "(99) 99999-9999"} unmask={true} />
                                        {submitted && !aluno.Telefone && (
                                        <small className="p-invalid">Telefone obrigatório</small>
                                    )}
                                </div>
                            </div>
                            <div className="field">
                                <label htmlFor="email">Email</label>
                                <InputText id="email" value={aluno.Email} required onChange={(e) => setAluno({...aluno, Email: e.target.value})} className={classNames({ 'p-invalid': submitted && !aluno.Email })} />
                                {submitted && !aluno.Email && <small className="p-invalid">Email obrigatório</small>}
                            </div>
                            <div className="field">
                                <label htmlFor="observacao">Nome</label>
                                <InputText id="observacao" value={aluno.Descricao} required autoFocus onChange={(e) => setAluno({...aluno, Descricao: e.target.value})} className={classNames({ 'p-invalid': submitted && !aluno.Descricao })} />
                                {submitted && !aluno.Descricao && <small className="p-invalid">{'Descricao obrigatório'}</small>}
                            </div>

                        </Dialog>

                        <Dialog visible={deleteAlunoDialog} style={{ width: '450px' }} header={'Confirmar'} modal footer={deleteAlunoDialogFooter} onHide={hideDeleteAlunoDialog}>
                            <div className="flex align-items-center justify-content-center">
                                <i className="pi pi-exclamation-triangle mr-3" style={{ fontSize: '2rem' }} />
                                {aluno && (
                                    <span>
                                        <b>Deseja excluir?</b>
                                    </span>
                                )}
                            </div>
                        </Dialog>

                        <BotaoVoltarMenu />


                    </div>
                </div>
            </div>

        </>
    );
}
