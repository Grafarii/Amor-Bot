# FofoqueiBot

Um bot de fofoca anônima para Discord.

## Funcionalidades
- `!fofocar <mensagem>`: Envia fofoca anônima com reações 👍 💩 ❗
- `!topfofocas [n]`: Lista as n fofocas mais votadas
- `!moderar <message_id>`: (Admin) Remove fofoca ofensiva

## Deploy no Railway
1. Crie projeto no Railway
2. Adicione plugin **PostgreSQL**
3. Configure secrets:
   - `DISCORD_TOKEN`
   - `DATABASE_URL`
4. Faça deploy (GitHub ou CLI)

Pronto!
