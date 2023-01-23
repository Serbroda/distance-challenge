<script lang="ts">
</script>

<main class="container">
    <br/>

    <article class="grid">
        <div>
            <hgroup>
                <h1>Sign up</h1>
                <h2>Create a new distance challenge account</h2>
            </hgroup>
            <form>
                <input
                        type="text"
                        name="username"
                        placeholder="Username"
                        aria-label="Username"
                        autocomplete="username"
                        required
                />
                <input
                        type="password"
                        name="password"
                        placeholder="Password"
                        aria-label="Password"
                        autocomplete="current-password"
                        required
                />
                <input
                        type="password"
                        name="passwordRepeat"
                        placeholder="Password (repeat)"
                        aria-label="Password (repeat)"
                        autocomplete="current-password"
                        required
                />

                <div class="text-center">
                    Already have an account? <a href="/#/login">Sign in</a>
                </div>
                <br/>
                <button type="submit" class="contrast" onclick="event.preventDefault()">Register</button>
            </form>
        </div>
        <div id="hero">
        </div>
    </article>

    <footer class="container-fluid">
        <small
        >Built with <a href="https://picocss.com" class="secondary">Pico</a>
            •
            <a href="https://github.com/picocss/examples/tree/master/sign-in/" class="secondary">Source code</a></small
        >
    </footer>
</main>

<style>
    article {
        padding: 0;
        overflow: hidden;
    }

    article div:nth-of-type(1) {
        padding: 1.5rem;
    }

    article div:nth-of-type(2) {
        display: none;
        background-image: url("https://cdn.midjourney.com/43fff174-005b-499f-a46b-ced0eab370ab/grid_0.png");
        background-position: center;
        background-size: cover;
    }

    .grid > div:nth-of-type(2) {
        display: block;
    }
</style>
