import { useContext } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { AuthContext } from "../context/auth";
import styles from './privateRoutes.module.scss';

export const PrivateRoute = () => {
  const { signed } = useContext(AuthContext);
  return signed ? 


        <div className={styles.container}>
          <div className={styles['container-outlet']}>
            <Outlet />
          </div>
        </div>
  
  : <Navigate to="/login" replace/>;
};