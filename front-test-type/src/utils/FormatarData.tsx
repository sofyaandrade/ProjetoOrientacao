export const formatarTempoGormParaTempoPadrao = (value: string | null) => {
    if(value){
        return new Date(value).toLocaleDateString('pt-br', {
            day: '2-digit',
            month: '2-digit',
            year: 'numeric',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit'
        });
    }
}

export const formatarTempoGormParaTempoData = (value: string | null) => {
    if(value){
        return new Date(value).toLocaleDateString('pt-br')
    }
}