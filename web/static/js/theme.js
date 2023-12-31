
const themeSwitcher = document.getElementById('theme-switcher');
const themeIcon = document.getElementById("path");

if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    document.documentElement.classList.add('dark');
    themeSwitcher.setAttribute('data-theme', "light");
    themeIcon.setAttribute('d', getPathForTheme("dark"))
} else {
    document.documentElement.classList.remove('dark');
    themeSwitcher.setAttribute('data-theme', "dark");
    themeIcon.setAttribute('d', getPathForTheme("light"))
}

function setDarkTheme() {
    document.documentElement.classList.add("dark");
    localStorage.theme = "dark";
    themeSwitcher.setAttribute('data-theme', "light");
    themeIcon.setAttribute('d', getPathForTheme("dark"))
}

function setLightTheme() {
    document.documentElement.classList.remove("dark");
    localStorage.theme = "light";
    themeSwitcher.setAttribute('data-theme', "dark");
    themeIcon.setAttribute('d', getPathForTheme("light"))
}

function onThemeSwitcherItemClick(event) {
    const theme = event.target.dataset.theme;
   if (theme === "dark") {
        setDarkTheme();
    } else {
        setLightTheme();
    }
}

function getPathForTheme(theme) {
    return theme === 'light' ?
        'M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z' :
        'M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69a.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z';
}

const themeSwitcherItems = document.querySelectorAll("#theme-switcher");
themeSwitcherItems.forEach((item) => {
    item.addEventListener("click", onThemeSwitcherItemClick);
});

