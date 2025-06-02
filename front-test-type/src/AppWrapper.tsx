import React, { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { AuthProvider } from './context/auth';
import { AppRouter } from './Routes/routes';
import { Provider } from 'react-redux';
import store from './store/store';

const AppWrapper = (props: any) => {
    const [colorScheme, setColorScheme] = useState('light');
    const [theme, setTheme] = useState('teal');
    const [componentTheme, setComponentTheme] = useState('teal');

    let location = useLocation();

    useEffect(() => {
        window.scrollTo(0, 0);
        onColorSchemeChange(colorScheme)
    }, [location]);

    const onColorSchemeChange = (scheme: string) => {
        changeStyleSheetUrl('layout-css', 'layout-teal.css', 1);
        changeStyleSheetUrl('theme-css', 'theme-teal.css', 1);
        setColorScheme(scheme);
    };

    const changeStyleSheetUrl = (id: any, value: any, from: any) => {
        const element = document.getElementById(id) as HTMLInputElement;
        const urlTokens = (element.getAttribute('href') as String).split('/');

        if (from === 1) {
            // which function invoked this function - change scheme
            urlTokens[urlTokens.length - 1] = value;
        } else if (from === 2) {
            // which function invoked this function - change color
            urlTokens[urlTokens.length - 2] = value;
        }

        const newURL = urlTokens.join('/');

        replaceLink(element, newURL);
    };

    const onMenuThemeChange = (theme: string) => {
        const layoutLink = document.getElementById('layout-css');
        const href = 'assets/layout/css/' + theme + '/layout-teal.css';

        replaceLink(layoutLink, href);
        setTheme(theme);
    };

    const onComponentThemeChange = (theme: string) => {
        const themeLink = document.getElementById('theme-css');
        const href = 'assets/theme/' + theme + '/theme-teal.css';

        replaceLink(themeLink, href);
        setComponentTheme(theme);
    };

    const replaceLink = (linkElement: any, href: string, callback?: any) => {
        const id = linkElement.getAttribute('id');
        const cloneLinkElement = linkElement.cloneNode(true);

        cloneLinkElement.setAttribute('href', href);
        cloneLinkElement.setAttribute('id', id + '-clone');

        linkElement.parentNode.insertBefore(cloneLinkElement, linkElement.nextSibling);

        cloneLinkElement.addEventListener('load', () => {
            linkElement.remove();
            const _linkElement = document.getElementById(id);
            _linkElement && _linkElement.remove();
            cloneLinkElement.setAttribute('id', id);

            if (callback) {
                callback();
            }
        });
    };

    return (
    <Provider store={store}>
        <AuthProvider>
            <AppRouter colorScheme={colorScheme} onColorSchemeChange={onColorSchemeChange} componentTheme={componentTheme} onComponentThemeChange={onComponentThemeChange} theme={theme} onMenuThemeChange={onMenuThemeChange} />
        </AuthProvider>
    </Provider>
    );
};

export default AppWrapper;
