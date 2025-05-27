import { Navigate, Outlet } from "react-router-dom";
import jwt_decode from "jwt-decode";

interface props {
    allowedRoles: string[];
}

const RequireAuth = ({ allowedRoles }: props) => {
    let token = {
        Role: ""
    }
    let role = "";
    let accessToken = sessionStorage.getItem("AuthAccessToken");
    if (accessToken !== null) {
        token = jwt_decode(accessToken);
        if (token && token.Role) {
            role = token.Role;
        }
    }

    let usuario = '';
    if (role != null) {
        usuario = role;
    }
    return (
        allowedRoles.includes(usuario)
            ? <Outlet />
            : <Navigate to="/login" replace />
    );

}

export default RequireAuth;