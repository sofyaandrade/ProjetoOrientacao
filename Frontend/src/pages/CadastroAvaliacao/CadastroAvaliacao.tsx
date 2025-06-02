import { InputText } from "primereact/inputtext"
import { Toast } from "primereact/toast"
import { Accordion, AccordionTab } from "primereact/accordion";
import { Button } from "primereact/button";

import { Panel } from 'primereact/panel';
import { useEffect, useRef, useState } from "react";
import { useAppDispatch, useAppSelector } from "../../store/hooks";
import IAvaliacao from "../../interface/IAvaliacao/IAvaliacao";
import { addAvaliacao, getAvaliacao, updateAvaliacao } from "../../service/avaliacaoService";
import { getAlunos } from "../../service/alunoService";
import IAluno from "../../interface/IAluno/IAluno";
import { useSelector } from "react-redux";
import { Fieldset } from 'primereact/fieldset';
import { Dropdown } from "primereact/dropdown";
import { InputNumber } from 'primereact/inputnumber';


export const CadastroAvaliacao = (props: any) => {
    const dispatch = useAppDispatch()
    const toast = useRef<any>(null);

    let avaliacaoVazia: IAvaliacao
    avaliacaoVazia = {
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

    const [avaliacao, setAvaliacao] = useState<IAvaliacao>(avaliacaoVazia)
    const [listaAlunos, setListaAlunos] = useState<IAluno[]>();
    const [submitted, setSubmitted] = useState<boolean>(false)
    const [idAluno, setIdAluno] = useState<IAvaliacao>()



    useEffect(() => {
        dispatch(getAvaliacao()).then((data) => {setAvaliacao(data.payload)})
        dispatch(getAlunos()).then((data) => setListaAlunos(data.payload))
    }, [])


    const salvarAvaliacao = () => {
        setSubmitted(true)
            dispatch(addAvaliacao(avaliacao))
            toast.current.show({ severity: 'success', summary: 'Sucesso', life: 3000 });
        setSubmitted(false)
    }

    const onAlunoChange = (e: any) => {
        setIdAluno(e.value)
        setAvaliacao({ ...avaliacao, AlunoID: Number(e.value) }) 
    }

    return (
        <>
            <h2></h2>
            <div className="col-12 flex mb-1 justify-content-center">
                <div className="w-8">
                    <div className="bg-white p-4">
                        <div >
                            <h5>Aluno</h5>
                            <div>
                                <Dropdown value={idAluno} options={listaAlunos} optionValue="ID" optionLabel="Descricao" onChange={(e) => onAlunoChange(e)} />
                            </div>
                            <div>

                            </div>
                        </div>
                    </div>

                    <Toast ref={toast} />


                    <div className="p-4 bg-sky-200 rounded-md">
                        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">

                            <div className="bg-white p-4 rounded shadow">
                                <h5 className="font-bold text-center mb-2">Tórax</h5>
                                <div className="field col-12 md:col-2">
                                    <span className="p-float-label">
                                        <InputNumber inputId="inspirado" value={avaliacao.ToraxInspirado} onChange={(e) => { setAvaliacao({ ...avaliacao, ToraxInspirado: Number(e.value) }) }} />
                                        <label htmlFor="inspirado">Inspirado</label>
                                    </span>
                                </div>   
                                <div className="field col-12 md:col-2">
                                    <span className="p-float-label">
                                        <InputNumber inputId="normal" value={avaliacao.ToraxNormal} onChange={(e) => { setAvaliacao({ ...avaliacao, ToraxNormal: Number(e.value) }) }} />
                                        <label htmlFor="normal">Normal</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-2">
                                    <span className="p-float-label">
                                        <InputNumber inputId="expirado" value={avaliacao.ToraxExpirado} onChange={(e) => { setAvaliacao({ ...avaliacao, ToraxExpirado: Number(e.value) }) }} />
                                        <label htmlFor="expirado">Expirado</label>
                                    </span>
                                </div> 
                            </div>

                            <div className="bg-white p-4 rounded shadow">
                                <h5 className="font-bold text-center mb-2">Tronco</h5>
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="cintura" value={avaliacao.Cintura} onChange={(e) => { setAvaliacao({ ...avaliacao, Cintura: Number(e.value) }) }} />
                                        <label htmlFor="cintura">Cintura</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="abdomen" value={avaliacao.Abdomen} onChange={(e) => { setAvaliacao({ ...avaliacao, Abdomen: Number(e.value) }) }} />
                                        <label htmlFor="abdomen">Abdomen</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="quadril" value={avaliacao.Quadril} onChange={(e) => { setAvaliacao({ ...avaliacao, Quadril: Number(e.value) }) }} />
                                        <label htmlFor="quadril">Quadril</label>
                                    </span>
                                </div> 
                            </div>

                            <div className="bg-white p-4 rounded shadow">
                                <h5 className="font-bold text-center mb-2">Dobras Cutâneas Masc</h5>
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="torax" value={avaliacao.Torax} onChange={(e) => { setAvaliacao({ ...avaliacao, Torax: Number(e.value) }) }} />
                                        <label htmlFor="torax">Torax</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="coxa" value={avaliacao.Coxa} onChange={(e) => { setAvaliacao({ ...avaliacao, Coxa: Number(e.value) }) }} />
                                        <label htmlFor="coxa">Coxa</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="abdomenMasc" value={avaliacao.AbdomenMasc} onChange={(e) => { setAvaliacao({ ...avaliacao, AbdomenMasc: Number(e.value) }) }} />
                                        <label htmlFor="abdomenMasc">Abdomen Masc</label>
                                    </span>
                                </div> 
                            </div>

                            <div className="bg-white p-4 rounded shadow">
                                <h5 className="font-bold text-center mb-2">Dobras Cutâneas Fem</h5>
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="triceps" value={avaliacao.Triceps} onChange={(e) => { setAvaliacao({ ...avaliacao, Triceps: Number(e.value) }) }} />
                                        <label htmlFor="triceps">Triceps</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="supraIliaca" value={avaliacao.SupraIliaca} onChange={(e) => { setAvaliacao({ ...avaliacao, SupraIliaca: Number(e.value) }) }} />
                                        <label htmlFor="supraIliaca">Supra-Ilíaca</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="coxa" value={avaliacao.Coxa} onChange={(e) => { setAvaliacao({ ...avaliacao, Coxa: Number(e.value) }) }} />
                                        <label htmlFor="coxa">Coxa</label>
                                    </span>
                                </div> 
                            </div>

                            <div className="bg-white p-4 rounded shadow col-span-2">
                                <h5 className="font-bold text-center mb-2">AnteBraço</h5>
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="anteDireito" value={avaliacao.AnteBracoDireito} onChange={(e) => { setAvaliacao({ ...avaliacao, AnteBracoDireito: Number(e.value) }) }} />
                                        <label htmlFor="anteDireito">Direito</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="anteEsquerdo" value={avaliacao.AnteBracoEsquerdo} onChange={(e) => { setAvaliacao({ ...avaliacao, AnteBracoEsquerdo: Number(e.value) }) }} />
                                        <label htmlFor="anteEsquerdo">Esquerdo</label>
                                    </span>
                                </div> 
                            </div>

                            <div className="bg-white p-4 rounded shadow col-span-2">
                                <h5 className="font-bold text-center mb-2">Braços Contraídos</h5>
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="contraidoDireiro" value={avaliacao.BracoContraidoDireito} onChange={(e) => { setAvaliacao({ ...avaliacao, BracoContraidoDireito: Number(e.value) }) }} />
                                        <label htmlFor="contraidoDireiro">Direito</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="contraidoEsquerdo" value={avaliacao.BracoContraidoEsquerdo} onChange={(e) => { setAvaliacao({ ...avaliacao, BracoContraidoEsquerdo: Number(e.value) }) }} />
                                        <label htmlFor="contraidoEsquerdo">Esquerdo</label>
                                    </span>
                                </div> 
                            </div>

                            <div className="bg-white p-4 rounded shadow col-span-2">
                                <h5 className="font-bold text-center mb-2">Braços Relaxados</h5>
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="relaxadoDireito" value={avaliacao.BracoRelaxadoDireito} onChange={(e) => { setAvaliacao({ ...avaliacao, BracoRelaxadoDireito: Number(e.value) }) }} />
                                        <label htmlFor="relaxadoDireito">Direito</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="relaxadoEsquerdo" value={avaliacao.BracoRelaxadoEsquerdo} onChange={(e) => { setAvaliacao({ ...avaliacao, BracoRelaxadoEsquerdo: Number(e.value) }) }} />
                                        <label htmlFor="relaxadoEsquerdo">Esquerdo</label>
                                    </span>
                                </div>   
                            </div>

                            <div className="bg-white p-4 rounded shadow col-span-2 ">
                                <h5 className="font-bold text-center mb-2">Coxas</h5>
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="coxaDireita" value={avaliacao.CoxaDireita} onChange={(e) => { setAvaliacao({ ...avaliacao, CoxaDireita: Number(e.value) }) }} />
                                        <label htmlFor="coxaDireita">Direita</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="coxaEsquerda" value={avaliacao.CoxaEsquerda} onChange={(e) => { setAvaliacao({ ...avaliacao, CoxaEsquerda: Number(e.value) }) }} />
                                        <label htmlFor="coxaEsquerda">Esquerda</label>
                                    </span>
                                </div> 
                            </div>

                            <div className="bg-white p-4 rounded shadow col-span-2">
                                <h5 className="font-bold text-center mb-2">Pernas</h5>
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="pernaDireita" value={avaliacao.PernaDireita} onChange={(e) => { setAvaliacao({ ...avaliacao, PernaDireita: Number(e.value) }) }} />
                                        <label htmlFor="pernaDireita">Direita</label>
                                    </span>
                                </div> 
                                <div className="field col-12 md:col-4">
                                    <span className="p-float-label">
                                        <InputNumber inputId="pernaEsquerda" value={avaliacao.PernaEsquerda} onChange={(e) => { setAvaliacao({ ...avaliacao, PernaEsquerda: Number(e.value) }) }} />
                                        <label htmlFor="pernaEsquerda">Esquerda</label>
                                    </span>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
            </div>


            <div className="col-12 flex justify-content-center">
            <Button 
                label={'Salvar'} onClick={() => { salvarAvaliacao() }} />
                {submitted}
            </div>
        </>
    )
}

function setListaAlunos(payload: any): any {
    throw new Error("Function not implemented.");
}
