#!/usr/bin/python3

from PyQt5.QtCore import QSize, Qt
from PyQt5.QtWidgets import *
from PyQt5.QtGui import QPixmap

import sys

class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()

        self.setWindowTitle("[Data Privacy] Can Data Be Queer?")
        self.setFixedSize(QSize(800, 600))
        # self.setStyleSheet("background-color: white;")

        self.label = QLabel(self)
        self.pixmap = QPixmap('splash.png').scaled(400, 400)

        self.label.setPixmap(self.pixmap)

        self.label.resize(self.pixmap.width(), self.pixmap.height())
        self.label.move(200, 0)

        self.client_widgets()


    def client_widgets(self):

        self.submit_label = QLabel("your name:", self)
        self.submit_label.move(200, 440)

        self.submit_textbox = QLineEdit(self)
        self.submit_textbox.move(300, 440)
        self.submit_textbox.resize(300, 30)

        self.question_label = QLabel("do you identify with the lgbtq+ community?", self)
        self.question_label.move(200, 500)
        self.question_label.resize(270, 30)

        self.yes_button = QPushButton("yes", self)
        self.yes_button.move(200, 550)

        self.no_button = QPushButton("no", self)
        self.no_button.move(300, 550)



app = QApplication(sys.argv)

window = MainWindow()
window.show()

app.exec()
