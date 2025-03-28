if (!("theme" in localStorage)) {
    // sets theme based on media preference
    localStorage.theme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
}
document.documentElement.classList.toggle(
    localStorage.theme, true,
);
