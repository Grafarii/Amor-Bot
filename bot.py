import discord
from discord.ext import commands
import os
from database.db import init_db
import asyncio

intents = discord.Intents.default()
intents.messages = True
intents.message_content = True
intents.reactions = True

bot = commands.Bot(command_prefix='!', intents=intents)

@bot.event
async def on_ready():
    print(f"Logged in as {bot.user}")
    await init_db()

# Load extensions
extensions = [
    'commands.fofoca',
    'commands.ranking',
    'commands.admin'
]

for ext in extensions:
    bot.load_extension(ext)

# Reaction tracking
@bot.event
async def on_raw_reaction_add(payload: discord.RawReactionActionEvent):
    if payload.member and payload.member.bot:
        return
    from database.db import AsyncSessionLocal
    from database.models import Fofoca
    async with AsyncSessionLocal() as session:
        fofoca = await session.get(Fofoca, payload.message_id)
        if not fofoca:
            return
        if str(payload.emoji) == '👍':
            fofoca.upvotes += 1
        elif str(payload.emoji) == '💩':
            fofoca.downvotes += 1
        elif str(payload.emoji) == '❗':
            fofoca.reports += 1
        await session.commit()

bot.run(os.getenv("DISCORD_TOKEN"))
