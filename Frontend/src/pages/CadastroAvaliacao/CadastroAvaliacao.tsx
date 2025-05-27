import { InputText } from "primereact/inputtext"
import { Toast } from "primereact/toast"
import { Accordion, AccordionTab } from "primereact/accordion";
import { Button } from "primereact/button";
import { useEffect, useRef, useState } from "react";
import { useAppDispatch } from "../../store/hooks";
import IAvaliacao from "../../interface/IAvaliacao/IAvaliacao";
import { addAvaliacao, getAvaliacao, updateAvaliacao } from "../../service/avaliacaoService";


export const CadastroAvaliacao = (props: any) => {
    const dispatch = useAppDispatch()
    const toast = useRef<any>(null);

    let avaliacaoVazia: IAvaliacao
    avaliacaoVazia = {
        ID: 0,
        Abdomen: 0,
        abdomem_masc: 0,
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
        SupraLiaca: 0,
        Torax: 0,
        ToraxExpirado: 0,
        ToraxInspirado: 0,
        ToraxNormal: 0,
        Triceps: 0,
        AlunoID: 0,
    }

    const [avaliacao, setAvaliacao] = useState<IAvaliacao>(avaliacaoVazia)
    const [submitted, setSubmitted] = useState<boolean>(false)


    useEffect(() => {
        dispatch(getAvaliacao()).then((data) => {setAvaliacao(data.payload)})
    }, [])


    const salvarAvaliacao = () => {
        setSubmitted(true)
            dispatch(addAvaliacao(avaliacao))
            toast.current.show({ severity: 'success', summary: 'Sucesso', life: 3000 });
        setSubmitted(false)
    }

    return (
        <>
            <h2></h2>
                <div className="col-12 flex mb-1 justify-content-center">
                    <div className="w-8">
                        <Toast ref={toast} />
                        <Accordion activeIndex={0}>
                            <AccordionTab header=''>
                                <div className="flex flex-wrap flex-row gap-2">
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="webservice">Torax Inspirado</label>
                                        <InputText id="webservice" value={avaliacao?.ToraxInspirado} onChange={(e) => { setAvaliacao({ ...avaliacao, ToraxInspirado: Number(e.target.value) }) }} />
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="intervaloTags">Torax Normal</label>
                                        <InputText id="intervaloTgas" value={avaliacao?.ToraxNormal} onChange={(e) => { setAvaliacao({ ...avaliacao, ToraxNormal: Number(e.target.value) }) }} />             
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="tempoMaximo">Torax Expirado</label>
                                        <InputText id="tempoMaximo" value={avaliacao?.ToraxExpirado} onChange={(e) => { setAvaliacao({ ...avaliacao, ToraxExpirado: Number(e.target.value) }) }} />    
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="leituraAllen">Cintura</label>
                                        <InputText id="leituraAllen" value={avaliacao?.Cintura} onChange={(e) => { setAvaliacao({ ...avaliacao, Cintura: Number(e.target.value) }) }} />
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="comunicacaoAllen">Abdomen</label>
                                        <InputText id="comunicacaoAllen" value={avaliacao?.Abdomen} onChange={(e) => { setAvaliacao({ ...avaliacao, Abdomen: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex items-center gap-2">
                                        <label htmlFor="leituraSiemens" >Quadril</label>
                                        <InputText id="leituraSiemens"  value={avaliacao?.Quadril} onChange={(e) => { setAvaliacao({ ...avaliacao, Quadril: Number(e.target.value) }) }} />
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="leituraModbus">Ante Braço Direito</label>
                                        <InputText id="leituraModbus"  value={avaliacao?.AnteBracoDireito} onChange={(e) => { setAvaliacao({ ...avaliacao, AnteBracoDireito: Number(e.target.value) }) }} />
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="comunicacaoModbus">Ante braço Esquerdo</label>
                                        <InputText id="comunicacaoModbus"  value={avaliacao?.AnteBracoEsquerdo} onChange={(e) => { setAvaliacao({ ...avaliacao, AnteBracoEsquerdo: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">Braço Contraído Direito</label>
                                        <InputText id="itfPorta"  value={avaliacao?.BracoContraidoDireito} onChange={(e) => { setAvaliacao({ ...avaliacao, BracoContraidoDireito: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.BracoContraidoEsquerdo} onChange={(e) => { setAvaliacao({ ...avaliacao, BracoContraidoEsquerdo: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.BracoRelaxadoDireito} onChange={(e) => { setAvaliacao({ ...avaliacao, BracoRelaxadoDireito: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.BracoRelaxadoEsquerdo} onChange={(e) => { setAvaliacao({ ...avaliacao, BracoRelaxadoEsquerdo: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.CoxaDireita} onChange={(e) => { setAvaliacao({ ...avaliacao, CoxaDireita: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.CoxaEsquerda} onChange={(e) => { setAvaliacao({ ...avaliacao, CoxaEsquerda: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.PernaDireita} onChange={(e) => { setAvaliacao({ ...avaliacao, PernaDireita: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.PernaEsquerda} onChange={(e) => { setAvaliacao({ ...avaliacao, PernaEsquerda: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.Torax} onChange={(e) => { setAvaliacao({ ...avaliacao, Torax: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.abdomem_masc} onChange={(e) => { setAvaliacao({ ...avaliacao, abdomem_masc: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.SupraLiaca} onChange={(e) => { setAvaliacao({ ...avaliacao, SupraLiaca: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.Triceps} onChange={(e) => { setAvaliacao({ ...avaliacao, Triceps: Number(e.target.value) }) }}/>
                                    </div>
                                    <div className="flex align-items-center gap-2">
                                        <label htmlFor="itfPorta">   </label>
                                        <InputText id="itfPorta"  value={avaliacao?.Coxa} onChange={(e) => { setAvaliacao({ ...avaliacao, Coxa: Number(e.target.value) }) }}/>
                                    </div>
                                </div>
                            </AccordionTab>
                        </Accordion>
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