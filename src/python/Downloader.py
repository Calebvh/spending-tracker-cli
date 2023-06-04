import mintapi
import configparser
import os
import pandas as pd

#Reads the username and password from the config.ini file
def getUserInfo():
    # Get the absolute path to the current directory
    current_directory = os.path.dirname(os.path.abspath(__file__))

    # Get the absolute path to the parent directory
    parent_directory = os.path.dirname(current_directory)

    # Get the absolute path to the configuration file
    config_file_path = os.path.join(parent_directory, 'config', 'config.ini')
    print(config_file_path)

    config = configparser.ConfigParser()
    config.read(config_file_path)
    print(config.get('mintCredentials', 'username'))
    try:
        username = config['mintCredentials']['username']
        password = config['mintCredentials']['password']
        return username, password
    except:
        print("Error reading config.ini file. Please check that the file exists and is formatted correctly.")
        return None, None

    


def download_transactions():
    username, password = getUserInfo()
    if username == None or password == None:
        return
    mint = mintapi.Mint(username, password)
    mint.initiate_account_refresh()
    download_transactions = mint.get_transactions_data()

    for transaction in download_transactions:
        print(transaction)

    mint.close()

download_transactions()