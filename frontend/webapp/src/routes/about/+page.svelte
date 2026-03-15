<script lang="ts">
    import { TitleBar, HeroBar, Card } from 'svebulma';
</script>

<TitleBar breadcrumbs={['Admin', 'About']} />
<HeroBar title="About" />

<section class="section is-main-section">
    <Card title="About this project" icon="information-outline">
        <div class="content">
            <p>This is the <strong>Golang-SvelteKit Boilerplate</strong> — a fullstack starter kit combining:</p>
            <ul>
                <li><strong>SvelteKit</strong> — Modern frontend framework</li>
                <li><strong>Golang</strong> — Server with embedded frontend</li>
                <li><strong>Bulma</strong> — CSS framework</li>
                <li><strong>Admin One</strong> — Dashboard template</li>
            </ul>
            <p>The dashboard UI is adapted from the free <a href="https://github.com/vikdiesel/admin-one-bulma-dashboard" target="_blank">Admin One Bulma Dashboard</a> template (MIT License).</p>
        </div>
    </Card>
</section>
