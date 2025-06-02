import './AppFooter.css';

const AppFooter = (props: any) => {
    return (
        <div className="layout-footer">
            <div className="footer-logo-container">
                {/* <img id="footer-logo" className="simbolo-footer" alt="atlantis-layout" /> */}
                <span className="app-name">Sofya</span>
            </div>
            <span className="copyright">&#169; Sofya</span>
        </div>
    );
};

export default AppFooter;
