#!/bin/bash

# Цветовое оформление для красоты
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

show_menu() {
    clear
    echo -e "${BLUE}=========================================${NC}"
    echo -e "${GREEN}      ПАНЕЛЬ УПРАВЛЕНИЯ TODO-APP${NC}"
    echo -e "${BLUE}=========================================${NC}"
    echo -e "1) 🚀 Запустить Базу Данных (env-up)"
    echo -e "2) 🛑 Остановить Базу Данных (env-down)"
    echo -e "3) 🌉 Открыть порт БД для хоста (socat/portf)"
    echo -e "4) 🔒 Закрыть порт БД (portf-close)"
    echo -e "5) 📈 Накатить миграции (migrate-up)"
    echo -e "6) 📉 Откатить миграции (migrate-down)"
    echo -e "7) ⚠️ 🔥 ПОЛНАЯ ОЧИСТКА ОКРУЖЕНИЯ (cleanup)"
    echo -e "0) ❌ Выйти из меню"
    echo -e "${BLUE}=========================================${NC}"
    echo -n "Выберите действие [0-7]: "
}

while true; do
    show_menu
    read choice
    case $choice in
        1) echo -e "\n${YELLOW}Запуск БД...${NC}"; make env-up; read -p "Нажмите Enter для продолжения..." ;;
        2) echo -e "\n${YELLOW}Остановка БД...${NC}"; make env-down; read -p "Нажмите Enter для продолжения..." ;;
        3) echo -e "\n${YELLOW}Проброс портов через socat...${NC}"; make env-portf; read -p "Нажмите Enter для продолжения..." ;;
        4) echo -e "\n${YELLOW}Закрытие портов...${NC}"; make env-portf-cloase; read -p "Нажмите Enter для продолжения..." ;; # Опечатка как в твоем Makefile
        5) echo -e "\n${YELLOW}Запуск миграций UP...${NC}"; make migrate-up; read -p "Нажмите Enter для продолжения..." ;;
        6) echo -e "\n${YELLOW}Откат миграций DOWN...${NC}"; make migrate-down; read -p "Нажмите Enter для продолжения..." ;;
        7) echo -e "\n${RED}Внимание! Очистка...${NC}"; make env-cleanup; read -p "Нажмите Enter для продолжения..." ;;
        0) echo -e "\n${GREEN}Пока!${NC}"; exit 0 ;;
        *) echo -e "\n${RED}Неверный выбор!${NC}"; sleep 1 ;;
    esac
done
