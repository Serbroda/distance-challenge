<script lang="ts">
    import Router from 'svelte-spa-router'

    import Home from './routes/Home.svelte'
    import Login from './routes/Login.svelte'
    import NotFound from './routes/NotFound.svelte'
    import Register from './routes/Register.svelte'
    import {init as initTheme} from "./utils/dom.utils";
    import ThemeSwitcher from "./lib/ThemeSwitcher.svelte";

    const routes = {
        '/': Home,
        '/login': Login,
        '/Register': Register,
        '*': NotFound,
    };

</script>

<main>
    <Router {routes} />
    <ThemeSwitcher/>
</main>

<style>
</style>
