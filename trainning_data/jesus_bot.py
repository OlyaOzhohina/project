import logging
import os

from dotenv import load_dotenv
from aiogram import Bot, Dispatcher, executor, types

load_dotenv()

logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
                    level=logging.INFO)

bot = Bot(token=os.environ.get('TOKEN'))
dp = Dispatcher(bot)

@dp.message_handler(commands=['start', 'help'])
async def send_welcome(message: types.Message):
   
    await message.reply("Hi! I'm Jesus_bot!I love you")

@dp.message_handler()
async def echo(message: types.Message):

    await message.answer(message.text)

if __name__ == '__main__':
    executor.start_polling(dp, skip_updates=True)
