import { createContext } from "react";
import { useState } from "react";
import { useEffect } from "react";
import { Navigate } from "react-router-dom";
import { api } from "../service/api";
import { useAppDispatch } from '../store/hooks';
import { fetchUsuarioLogado } from "../store/reducers/usuarioLogado";
import jwt_decode from "jwt-decode";
export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const dispatch = useAppDispatch();

  useEffect(() => {
    const loadingStoreData = () => {
      const storageAccessToken = sessionStorage.getItem("AuthAccessToken");
      const storageRefreshToken = sessionStorage.getItem("AuthRefreshToken");

      if (storageAccessToken && storageRefreshToken) {
        setUser(storageAccessToken);
      }
    };
    loadingStoreData();
  }, []);

  const signIn = async ({ email, password }) => {
    try {
      const response = await api.post("/login/", { email, password });
      if (response.data.error) {
        alert(response.data.error);
      } else {
        setUser(response.data);
        sessionStorage.setItem("AuthAccessToken", response.data.AccessToken);
        sessionStorage.setItem("AuthRefreshToken", response.data.RefreshToken);

        const token = jwt_decode(response.data.AccessToken);
        sessionStorage.setItem("Id", token.Id);
        sessionStorage.setItem("Perfil", token.Role);

        dispatch(fetchUsuarioLogado());
      }
    } catch (error) {
      console.log(error);
      return error.response.status;
    }
  };

  const singOut = () => {
    sessionStorage.clear();
    setUser(null);
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        signIn,
        singOut,
        signed: !!user,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};