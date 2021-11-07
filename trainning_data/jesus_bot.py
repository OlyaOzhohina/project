import logging
import os

from dotenv import load_dotenv
from telegram.ext import Updater, CommandHandler,  MessageHandler, Filters

load_dotenv()

logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
                    level=logging.INFO)

updater = Updater(token=os.environ.get('TOKEN'), use_context=True)
dispatcher = updater.dispatcher


def start(update, context):
    context.bot.send_message(
        chat_id=update.effective_chat.id, text="I'm a Jesus bot!")


def unknown(update, context):
    context.bot.send_message(chat_id=update.effective_chat.id,
                             text="Sorry, I didn't understand that command.")


start_handler = CommandHandler('start', start)
unknown_handler = MessageHandler(Filters.command, unknown)

dispatcher.add_handler(start_handler)
dispatcher.add_handler(unknown_handler)

updater.start_polling()
updater.idle()
