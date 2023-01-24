<script lang="ts">
    import ThemeSwitcher from "../lib/ThemeSwitcher.svelte";
    import SortIcon from "../lib/SortIcon.svelte";
    import type {OrderBy} from "../models/sort";

    let data = [
        {
            name: "Danny",
            distance: 11.5,
            time: 24.2 * 4,
        },
        {
            name: "Sascha",
            distance: 23.5,
            time: 30.2 * 4,
        },
    ]

    let sortBy: OrderBy = {col: "name", ascending: true};

    $: sort = (column) => {

        if (sortBy.col == column) {
            sortBy.ascending = !sortBy.ascending
        } else {
            sortBy.col = column
            sortBy.ascending = true
        }

        // Modifier to sorting function for ascending or descending
        let sortModifier = (sortBy.ascending) ? 1 : -1;

        let sort = (a, b) =>
            (a[column] < b[column])
                ? -1 * sortModifier
                : (a[column] > b[column])
                    ? 1 * sortModifier
                    : 0;

        data = data.sort(sort);
    }
</script>

<main>
    <nav class="container-fluid">
        <ul>
            <li><a href="./" class="contrast brand" onclick="event.preventDefault()"><strong>dc.</strong></a></li>
        </ul>
        <ul>
            <li>
                <details role="list" dir="rtl">
                    <summary aria-haspopup="listbox" role="link" class="secondary">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                             stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                             class="feather feather-user">
                            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                            <circle cx="12" cy="7" r="4"></circle>
                        </svg>
                    </summary>
                    <ul role="listbox">
                        <li><a href="/#/login">Logout</a></li>
                    </ul>
                </details>
            </li>
        </ul>
    </nav>

    <main class="container">
        <h1>Home</h1>

        <table role="grid">
            <thead>
            <tr>
                <th scope="col" class="pointer" on:click={sort('name')}>Name {#if sortBy.col === "name"}<SortIcon ascending={sortBy.ascending}/>{/if}</th>
                <th scope="col" class="pointer" on:click={sort('distance')}>Distance {#if sortBy.col === "distance"}<SortIcon ascending={sortBy.ascending}/>{/if}</th>
                <th scope="col" class="pointer" on:click={sort('time')}>Time {#if sortBy.col === "time"}<SortIcon ascending={sortBy.ascending}/>{/if}</th>
            </tr>
            </thead>
            <tbody>
            {#each data as p}
                <tr>
                    <td>{p.name}</td>
                    <td>{p.distance}</td>
                    <td>{p.time}</td>
                </tr>
            {/each}

            </tbody>
            <!--<tfoot>
            <tr>
                <th scope="col">#</th>
                <td scope="col">Total</td>
                <td scope="col">Total</td>
                <td scope="col">Total</td>
                <td scope="col">Total</td>
                <td scope="col">Total</td>
            </tr>
            </tfoot>-->
        </table>
    </main>

    <footer class="container-fluid">
        <small
        >Built with <a href="https://picocss.com" class="secondary">Pico</a>
            •
            <a href="https://github.com/picocss/examples/tree/master/sign-in/" class="secondary">Source code</a></small
        >
    </footer>
</main>

<style>
    nav {
        box-shadow: 0 1px 0 rgba(115, 130, 140, 0.2);
        margin-bottom: 20px;
    }

    .pointer {
        cursor: pointer;
    }
</style>
