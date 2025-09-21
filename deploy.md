# 🚀 Развертывание SAA Research на DigitalOcean

## 📋 Быстрое развертывание

### На сервере 104.248.70.69:

```bash
# 1. Клонирование репозитория
cd /opt
git clone https://github.com/your-username/saa-research.git
cd saa-research

# 2. Остановка старых процессов
pkill -f "go run"
systemctl stop nginx

# 3. Запуск SAA Research
go run main.go &

# 4. Проверка
curl http://localhost/api/health
```

## 🌐 Доступ

- **Главная страница:** http://104.248.70.69
- **SAA Capital Models:** http://104.248.70.69/saa-capital-models
- **API:** http://104.248.70.69/api/health

## 🔄 Обновления

```bash
cd /opt/saa-research
git pull origin main
pkill -f "go run"
go run main.go &
```

## ✅ Проверка работы

```bash
# Статус процесса
ps aux | grep main.go

# API здоровье
curl http://localhost/api/health

# Данные новостей
curl http://localhost/api/news
```
