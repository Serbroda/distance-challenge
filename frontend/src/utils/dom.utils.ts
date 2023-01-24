export type Theme = "light" | "dark";

const getThemeLocalStorage = (): Theme | undefined => {
    return localStorage.getItem("theme") as Theme;
}

const getPreferesColorTheme = (): Theme | undefined => {
    if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
        return "dark";
    }
    return undefined;
}

const getTheme = (): Theme => {
    return getThemeLocalStorage() || getPreferesColorTheme() || "light"
}

const setTheme = (theme: Theme) => {
    const htmlElement = document.querySelector("html");
    htmlElement.dataset.theme = theme;
    localStorage.setItem("theme", theme);
}

const toggleTheme = () => {
    setTheme(getTheme() === "light" ? "dark" : "light");
};

const init = (): Theme => {
    const t = getTheme();
    setTheme(t);
    return t;
}

export {init, getTheme, setTheme, toggleTheme, getThemeLocalStorage, getPreferesColorTheme}
