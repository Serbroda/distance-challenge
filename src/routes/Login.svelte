<script lang="ts">
</script>

<main class="container">
    <article class="grid">
        <div>
            <hgroup>
                <h1>Sign in</h1>
                <h2>A minimalist layout for Login pages</h2>
            </hgroup>
            <form>
                <input
                    type="text"
                    name="login"
                    placeholder="Login"
                    aria-label="Login"
                    autocomplete="nickname"
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
                <fieldset>
                    <label for="remember">
                        <input type="checkbox" role="switch" id="remember" name="remember" />
                        Remember me
                    </label>
                </fieldset>
                <button type="submit" class="contrast" onclick="event.preventDefault()">Login</button>
            </form>
        </div>
        <div>
           <img alt="runners" id="art" src="https://media.discordapp.net/attachments/1060144354448134195/1066465929355071488/Serbroda_people_running_threw_the_nature_realistic_cinematic_li_20a14331-83ac-4dd8-b2fd-4fd5e675401e.png?width=671&height=671"/>
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
    #art {
        background-color: #374956;
        background-position: center;
        background-size: cover;
    }
</style>
