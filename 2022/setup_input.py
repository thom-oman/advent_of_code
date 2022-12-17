import sys
import os
import requests
from selenium import webdriver
from selenium.webdriver.firefox.options import Options as FirefoxOptions

day = sys.argv[1]
year = 2022
dir = "{}".format(day)

if not os.path.exists(dir):
    os.makedirs(dir)
else:
    print("already exists!")

os.chdir(dir)

url = "https://adventofcode.com/{}/day/{}/input".format(year, day)

# x = requests.get(url)
# f = open("inputs.txt", "a")
# f.write(x.text)
# f.close()

# options = FirefoxOptions()
# options.add_argument("--headless")
# options.add_argument("user-data-dir=/Users/toman/Library/Application\ Support/Firefox/Profiles/28e55o24.default")
# driver = webdriver.Firefox(options=options)
# driver.get(url)

# print(driver.page_source)
# driver.close()
