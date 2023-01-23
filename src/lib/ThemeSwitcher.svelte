<script lang="ts">
    import {getTheme, init, toggleTheme} from "../utils/dom.utils";

    let activeTheme = init();

    const toggle = () => {
        toggleTheme();
        activeTheme = getTheme();
    };
</script>

<div on:click={toggle}>

    {#if activeTheme === "dark"}
        <div class="theme-switcher dark">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                 stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                 class="feather feather-sun">
                <circle cx="12" cy="12" r="5"></circle>
                <line x1="12" y1="1" x2="12" y2="3"></line>
                <line x1="12" y1="21" x2="12" y2="23"></line>
                <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
                <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
                <line x1="1" y1="12" x2="3" y2="12"></line>
                <line x1="21" y1="12" x2="23" y2="12"></line>
                <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
                <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
            </svg>
        </div>
    {:else }
        <div class="theme-switcher light">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                 stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                 class="feather feather-moon">
                <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
            </svg>
        </div>
    {/if}
</div>

<style>
    .theme-switcher {
        position: absolute;
        right: 25px;
        bottom: 25px;
        cursor: pointer;
        background-color: black;
        border-radius: 50%;
        width: 45px;
        height: 45px;
        color: white;
        display: flex;
        align-items: center;
    }

    .theme-switcher svg {
        margin: 0 auto;
        display: block;
    }

    .theme-switcher.dark {
        border: 1px solid #3d3d3d;
    }

    .theme-switcher.light:hover {
        color: #3398ea;
    }

    .theme-switcher.dark:hover {
        color: #f8bb21;
    }
</style>
