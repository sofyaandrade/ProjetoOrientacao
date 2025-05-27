import { ReactElement, JSXElementConstructor, ReactFragment, ReactPortal } from "react";
import "./styles.css";
export const LayoutComponents = (props: { children: string | number | boolean | ReactElement<any, string | JSXElementConstructor<any>> | ReactFragment | ReactPortal | null | undefined; }) => {
  return (
    <div className="container">
      <div className="container-login">
        <div className="wrap-login">{props.children}</div>
      </div>
    </div>
  );
};