import { Provider } from 'react-redux';
import { Routes, Route } from 'react-router-dom';
import store from '../store/store';
import RequireAuth from '../components/RequireAuth/RequireAuth';
import App from '../App';
import { Login } from '../pages/Login/Login';
import { isStandalone } from '../utils/envConfig';
import { NotFound } from '../pages/NotFound';
import { CadastroAvaliacao } from '../pages/CadastroAvaliacao/CadastroAvaliacao';
import { CadastroAluno } from '../pages/CadastroAluno/CadastroAluno';
import MainMenu from '../MainMenu';
import { CadastroUsuario } from '../pages/CadastroInstrutor/CadastroInstrutor';

export const AppRouter = (props: any) => {

    const ADMINISTRADOR = 'ADMINISTRADOR';
    const CLIENTE = 'CLIENTE';

    const standaloneAppRoutes =
        <>
            <Route element={<RequireAuth allowedRoles={[ADMINISTRADOR, CLIENTE]} />}>
                <Route index element={<MainMenu  />} />
                <Route path="*" element={<NotFound />} />
            </Route>
        </>

    const fullAppRoutes =
        <>
            {/* <Route element={<RequireAuth allowedRoles={[ADMINISTRADOR, CLIENTE]} />}>
                <Route index element={<Home colorScheme={props.colorScheme} />} />
            </Route> */}

            <Route element={<RequireAuth allowedRoles={[ADMINISTRADOR, CLIENTE]} />}>
                <Route path='menu-principal' element={<MainMenu />} />
            </Route>

            <Route element={<RequireAuth allowedRoles={[ADMINISTRADOR, CLIENTE]} />}>
                <Route path='cadastro-aluno' element={<CadastroAluno />} />
            </Route>

            <Route element={<RequireAuth allowedRoles={[ADMINISTRADOR, CLIENTE]} />}>
                <Route path='cadastro-instrutor' element={<CadastroUsuario />} />
            </Route>

            <Route element={<RequireAuth allowedRoles={[ADMINISTRADOR, CLIENTE]} />}>
                <Route path='cadastro-avaliacao' element={<CadastroAvaliacao />} />
            </Route>
        </>

    const configRoutes = () => {
        return isStandalone ? standaloneAppRoutes : fullAppRoutes
    }

    return (
            <Provider store={store}>
                <Routes>
                    <Route path='/login' element={<Login />} />

                    <Route path='*' element={<App colorScheme={props.colorScheme} onColorSchemeChange={props.onColorSchemeChange} componentTheme={props.componentTheme} onComponentThemeChange={props.onComponentThemeChange} theme={props.theme} onMenuThemeChange={props.onMenuThemeChange} />} >
                        {configRoutes()}
                    </Route>
                </Routes>
            </Provider>
    );
}