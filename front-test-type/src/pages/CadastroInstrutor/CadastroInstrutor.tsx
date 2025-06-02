import { InputText } from "primereact/inputtext"
import { Toast } from "primereact/toast"
import { Accordion, AccordionTab } from "primereact/accordion";
import { Button } from "primereact/button";
import { useEffect, useRef, useState } from "react";
import { useAppDispatch } from "../../store/hooks";
import IAvaliacao from "../../interface/IAvaliacao/IAvaliacao";
import { addAvaliacao, getAvaliacao, updateAvaliacao } from "../../service/avaliacaoService";
//import IInstrutor from "../../interface/IInstrutor/IInstrutor";
//import { addInstrutor, getInstrutor, updateInstrutor } from "../../service/clienteService";

// export const CadastroInstrutor = (props: any) => {
//     const dispatch = useAppDispatch()
//     const toast = useRef<any>(null);

//     let avaliacaoVazia: IAvaliacao
//     avaliacaoVazia = {
//         ID: 0,
//         Abdomen: 0,
//         abdomem_masc: 0,
//         AnteBracoDireito: 0,
//         AnteBracoEsquerdo: 0,
//         BracoContraidoDireito: 0,
//         BracoContraidoEsquerdo: 0,
//         BracoRelaxadoDireito: 0,
//         BracoRelaxadoEsquerdo: 0,
//         Cintura: 0,
//         Coxa: 0,
//         CoxaDireita: 0,
//         CoxaEsquerda: 0,
//         PernaDireita: 0,
//         PernaEsquerda: 0,
//         Quadril: 0,
//         SupraLiaca: 0,
//         Torax: 0,
//         ToraxExpirado: 0,
//         ToraxInspirado: 0,
//         ToraxNormal: 0,
//         Triceps: 0,
//         UsuarioID: 0,
//     }

//     let clienteVazio: IInstrutor
//     clienteVazio = {
//         ID: 0,
//         Nome: "",
//         Email: "",
//         Telefone: 0,
//         Descricao: "",
//         Avaliacao: {
//             ID: 0,
//             Abdomen: 0,
//             abdomem_masc: 0,
//             AnteBracoDireito: 0,
//             AnteBracoEsquerdo: 0,
//             BracoContraidoDireito: 0,
//             BracoContraidoEsquerdo: 0,
//             BracoRelaxadoDireito: 0,
//             BracoRelaxadoEsquerdo: 0,
//             Cintura: 0,
//             Coxa: 0,
//             CoxaDireita: 0,
//             CoxaEsquerda: 0,
//             PernaDireita: 0,
//             PernaEsquerda: 0,
//             Quadril: 0,
//             SupraLiaca: 0,
//             Torax: 0,
//             ToraxExpirado: 0,
//             ToraxInspirado: 0,
//             ToraxNormal: 0,
//             Triceps: 0,
//             UsuarioID: 0,
//         }
//     }
//     const [instrutor, setInstrutor] = useState<IInstrutor>(clienteVazio)
//     const [submitted, setSubmitted] = useState<boolean>(false)


//     useEffect(() => {

//     }, [])

//     const salvarInstrutor = () => {
//         setSubmitted(true)
//         dispatch(addInstrutor(cliente))
//         .unwrap()
//         .then(() => {
//           toast.current.show({ severity: 'success', summary: 'Sucesso', life: 3000 });
//         })
//         .catch((err) => {
//           console.error('Erro ao salvar cliente:', err);
//           toast.current.show({ severity: 'error', summary: 'Erro ao salvar cliente', detail: err.message, life: 5000 });
//         });
//     }

//     return (
//         <>
//             <h2></h2>
//                 <div className="col-12 flex mb-1 justify-content-center">
//                     <div className="w-8">
//                         <Toast ref={toast} />
//                         <Accordion activeIndex={0}>
//                             <AccordionTab header=''>
//                                 <div className="flex flex-wrap flex-row gap-2">
//                                     <div className="flex align-items-center gap-2">
//                                         <label htmlFor="webservice">Nome</label>
//                                         <InputText id="webservice" value={cliente?.Nome} onChange={(e) => { setInstrutor({ ...cliente, Nome: e.target.value }) }} />
//                                     </div>
//                                     <div className="flex align-items-center gap-2">
//                                         <label htmlFor="intervaloTags">Email</label>
//                                         <InputText id="intervaloTgas" value={cliente?.Email} onChange={(e) => { setInstrutor({ ...cliente, Email: e.target.value}) }} />             
//                                     </div>
//                                     <div className="flex align-items-center gap-2">
//                                         <label htmlFor="tempoMaximo">Telefone</label>
//                                         <InputText id="tempoMaximo" value={cliente?.Telefone} onChange={(e) => { setInstrutor({ ...cliente, Telefone: Number(e.target.value) }) }} />    
//                                     </div>
//                                     <div className="flex align-items-center gap-2">
//                                         <label htmlFor="leituraAllen">Descrição</label>
//                                         <InputText id="leituraAllen" value={cliente?.Descricao} onChange={(e) => { setInstrutor({ ...cliente, Descricao: e.target.value}) }} />
//                                     </div>
//                                 </div>
//                             </AccordionTab>
//                         </Accordion>
//                     </div>
//                 </div>


//             <div className="col-12 flex justify-content-center">
//             <Button 
//                 label={'Salvar'} onClick={() => { salvarInstrutor() }} />
//                 {submitted}
//             </div>
//         </>
//     )
// }