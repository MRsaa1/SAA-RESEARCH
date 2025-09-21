package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type NewsItem struct {
	Title      string `json:"title"`
	Sector     string `json:"sector"`
	Impact     int    `json:"impact"`
	Confidence int    `json:"confidence"`
	Direction  string `json:"direction"`
	Timestamp  string `json:"timestamp"`
}

var newsData = []NewsItem{
	{"Bitcoin достигает нового максимума $67,000", "КРИПТОВАЛЮТЫ", 95, 90, "positive", "2025-09-21T14:30:00Z"},
	{"ФРС сохраняет процентную ставку на уровне 5.25%", "КАЗНАЧЕЙСТВО", 88, 95, "neutral", "2025-09-21T14:25:00Z"},
	{"NVIDIA представляет новые AI-чипы следующего поколения", "ТЕХНОЛОГИИ", 82, 89, "positive", "2025-09-21T14:20:00Z"},
	{"Цены на нефть Brent превышают $90 за баррель", "ЭНЕРГЕТИКА", 75, 85, "negative", "2025-09-21T14:15:00Z"},
	{"Pfizer получает одобрение FDA на новый препарат", "ЗДРАВООХРАНЕНИЕ", 78, 91, "positive", "2025-09-21T14:10:00Z"},
	{"Apple анонсирует iPhone 16 с улучшенными камерами", "ТЕХНОЛОГИИ", 85, 88, "positive", "2025-09-21T14:05:00Z"},
	{"Ethereum успешно проводит обновление Dencun", "КРИПТОВАЛЮТЫ", 80, 87, "positive", "2025-09-21T14:00:00Z"},
	{"Tesla увеличивает производство электромобилей на 25%", "АВТОМОБИЛИ", 77, 83, "positive", "2025-09-21T13:55:00Z"},
}

func main() {
	// Главная страница и SAA Capital Models
	http.HandleFunc("/", handleMainPage)
	http.HandleFunc("/saa-capital-models", handleMainPage)

	// API endpoints
	http.HandleFunc("/api/health", handleHealth)
	http.HandleFunc("/api/news", handleNews)
	http.HandleFunc("/api/refresh", handleRefresh)

	fmt.Println("🚀 SAA Research starting on port 80")
	fmt.Printf("🌐 Dashboard: http://104.248.70.69\n")
	fmt.Printf("📊 API Health: http://104.248.70.69/api/health\n")
	fmt.Printf("📰 News API: http://104.248.70.69/api/news\n")

	http.ListenAndServe(":80", nil)
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	html := `<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>SAA Alliance | Новостной аналитический портал</title>
<style>
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', system-ui, sans-serif;
    background: #000;
    color: #fff;
    line-height: 1.6;
}

.container {
    max-width: 1400px;
    margin: 0 auto;
    padding: 20px;
}

.main-header {
    background: linear-gradient(135deg, #1a1a1a 0%, #2d2d2d 100%);
    border: 2px solid #d4af37;
    border-radius: 15px;
    padding: 40px;
    text-align: center;
    margin-bottom: 20px;
    box-shadow: 0 8px 32px rgba(212, 175, 55, 0.1);
}

.main-header h1 {
    color: #d4af37;
    font-size: 2.8rem;
    font-weight: 700;
    margin-bottom: 15px;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.main-header p {
    color: #cccccc;
    font-size: 1.2rem;
    margin-bottom: 25px;
}

.language-switcher {
    display: flex;
    justify-content: center;
    gap: 15px;
}

.lang-btn {
    background: #3498db;
    color: white;
    border: none;
    padding: 10px 25px;
    border-radius: 25px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
}

.lang-btn.inactive {
    background: #555;
    color: #999;
}

.controls-panel {
    background: linear-gradient(135deg, #1a1a1a 0%, #2d2d2d 100%);
    border: 2px solid #d4af37;
    border-radius: 15px;
    padding: 30px;
    margin-bottom: 20px;
}

.filters-grid {
    display: grid;
    grid-template-columns: repeat(7, 1fr);
    gap: 20px;
    margin-bottom: 25px;
}

.filter-group {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.filter-group label {
    color: #cccccc;
    font-size: 0.9rem;
    font-weight: 600;
}

.filter-group select,
.filter-group input {
    background: #333;
    border: 1px solid #666;
    color: #fff;
    padding: 10px 12px;
    border-radius: 5px;
    font-size: 0.9rem;
}

.action-buttons {
    display: flex;
    justify-content: center;
    gap: 15px;
}

.action-btn {
    background: linear-gradient(135deg, #555 0%, #777 50%, #555 100%);
    color: #fff;
    border: none;
    padding: 12px 30px;
    border-radius: 8px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    text-transform: uppercase;
    letter-spacing: 1px;
}

.action-btn:hover {
    background: linear-gradient(135deg, #d4af37 0%, #f4d03f 100%);
    color: #000;
    transform: translateY(-2px);
    box-shadow: 0 5px 15px rgba(212, 175, 55, 0.3);
}

.action-btn.export {
    background: linear-gradient(135deg, #4a90e2 0%, #357abd 100%);
}

.stats-grid {
    display: grid;
    grid-template-columns: repeat(8, 1fr);
    gap: 15px;
    margin-bottom: 20px;
}

.stat-card {
    background: linear-gradient(135deg, #1a1a1a 0%, #2d2d2d 100%);
    border: 2px solid #d4af37;
    border-radius: 12px;
    padding: 25px 15px;
    text-align: center;
    transition: transform 0.2s ease;
}

.stat-card:hover {
    transform: translateY(-3px);
}

.stat-number {
    font-size: 2.5rem;
    font-weight: 700;
    color: #d4af37;
    margin-bottom: 8px;
}

.stat-label {
    color: #cccccc;
    font-size: 0.8rem;
    text-transform: uppercase;
    letter-spacing: 1px;
    line-height: 1.2;
}

.signals-panel {
    background: linear-gradient(135deg, #1a1a1a 0%, #2d2d2d 100%);
    border: 2px solid #d4af37;
    border-radius: 15px;
    overflow: hidden;
}

.signals-header {
    background: #333;
    padding: 20px 30px;
    border-bottom: 2px solid #d4af37;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.signals-title {
    color: #d4af37;
    font-size: 1.3rem;
    font-weight: 600;
    display: flex;
    align-items: center;
    gap: 10px;
}

.signals-content {
    padding: 30px;
    min-height: 300px;
}

.signal-item {
    background: #2a2a2a;
    margin: 15px 0;
    padding: 20px;
    border-radius: 8px;
    border-left: 5px solid #d4af37;
    transition: all 0.2s ease;
}

.signal-item:hover {
    transform: translateX(5px);
    box-shadow: 0 5px 15px rgba(212, 175, 55, 0.2);
}

.signal-title {
    color: #ffffff;
    font-weight: 600;
    margin-bottom: 10px;
    font-size: 1.1rem;
    line-height: 1.4;
}

.signal-meta {
    color: #999;
    font-size: 0.95rem;
    display: flex;
    gap: 20px;
    flex-wrap: wrap;
}

.loading-text {
    text-align: center;
    color: #999;
    font-style: italic;
    font-size: 1.1rem;
    padding: 40px;
}

@media (max-width: 1200px) {
    .filters-grid { grid-template-columns: repeat(4, 1fr); }
    .stats-grid { grid-template-columns: repeat(4, 1fr); }
}

@media (max-width: 768px) {
    .filters-grid { grid-template-columns: repeat(2, 1fr); }
    .stats-grid { grid-template-columns: repeat(2, 1fr); }
    .main-header h1 { font-size: 2.2rem; }
}
</style>
</head>
<body>
<div class="container">
    <div class="main-header">
        <h1>SAA Alliance | Новостной аналитический портал</h1>
        <p>Профессиональная система аналитики</p>
        <div class="language-switcher">
            <button class="lang-btn">🇷🇺 Русский</button>
            <button class="lang-btn inactive">🇬🇧 English (скоро)</button>
        </div>
    </div>

    <div class="controls-panel">
        <div class="filters-grid">
            <div class="filter-group">
                <label>Сектор</label>
                <select><option>Все секторы</option></select>
            </div>
            <div class="filter-group">
                <label>Настроение рынка</label>
                <select><option>Все настроения</option></select>
            </div>
            <div class="filter-group">
                <label>Регион</label>
                <select><option>Все регионы</option></select>
            </div>
            <div class="filter-group">
                <label>Мин. влияние</label>
                <input type="number" value="40" min="0" max="100">
            </div>
            <div class="filter-group">
                <label>Мин. достоверность</label>
                <input type="number" value="0" min="0" max="100">
            </div>
            <div class="filter-group">
                <label>📅 Дата новостей</label>
                <input type="date" value="2025-09-19">
            </div>
            <div class="filter-group">
                <label>Поиск</label>
                <input type="text" placeholder="Поиск по новостям...">
            </div>
        </div>
        <div class="action-buttons">
            <button class="action-btn" onclick="loadSignals()">🔄 ЗАГРУЗИТЬ СИГНАЛЫ</button>
            <button class="action-btn export">📊 ЭКСПОРТ ДАННЫХ</button>
        </div>
    </div>

    <div class="stats-grid">
        <div class="stat-card">
            <div class="stat-number" id="total-signals">0</div>
            <div class="stat-label">ВСЕГО СИГНАЛОВ</div>
        </div>
        <div class="stat-card">
            <div class="stat-number" id="high-impact">0</div>
            <div class="stat-label">ВЫСОКОЕ ВЛИЯНИЕ (70+)</div>
        </div>
        <div class="stat-card">
            <div class="stat-number" id="medium-impact">0</div>
            <div class="stat-label">СРЕДНЕЕ ВЛИЯНИЕ</div>
        </div>
        <div class="stat-card">
            <div class="stat-number" id="avg-confidence">0%</div>
            <div class="stat-label">СР. ДОСТОВЕРНОСТЬ</div>
        </div>
        <div class="stat-card">
            <div class="stat-number" id="bull-signals">0</div>
            <div class="stat-label">🐃 БЫЧЬИ СИГНАЛЫ</div>
        </div>
        <div class="stat-card">
            <div class="stat-number" id="bear-signals">0</div>
            <div class="stat-label">🐻 МЕДВЕЖЬИ СИГНАЛЫ</div>
        </div>
        <div class="stat-card">
            <div class="stat-number" id="active-sectors">0</div>
            <div class="stat-label">АКТИВНЫХ СЕКТОРОВ</div>
        </div>
        <div class="stat-card">
            <div class="stat-number" id="regions">0</div>
            <div class="stat-label">РЕГИОНОВ</div>
        </div>
    </div>

    <div class="signals-panel">
        <div class="signals-header">
            <div class="signals-title">📊 Инвестиционные сигналы</div>
            <div style="color: #999;">Обновлено: <span id="update-time">19.09.2025, 14:30:06</span></div>
        </div>
        <div class="signals-content">
            <div id="signals-list" class="loading-text">Загрузка сигналов...</div>
        </div>
    </div>
</div>

<script>
function loadSignals() {
    document.getElementById('signals-list').innerHTML = '<div class="loading-text">Загрузка данных...</div>';
    
    fetch('/api/news')
    .then(response => response.json())
    .then(data => {
        let html = '';
        let highImpact = 0;
        let mediumImpact = 0;
        let bullSignals = 0;
        let bearSignals = 0;
        let sectors = new Set();
        let totalConfidence = 0;
        
        data.forEach(item => {
            sectors.add(item.sector);
            totalConfidence += item.confidence;
            
            if (item.impact >= 80) highImpact++;
            else if (item.impact >= 60) mediumImpact++;
            
            if (item.direction === 'positive') bullSignals++;
            if (item.direction === 'negative') bearSignals++;
            
            html += '<div class="signal-item">';
            html += '<div class="signal-title">' + item.title + '</div>';
            html += '<div class="signal-meta">';
            html += '<span>📊 ' + item.sector + '</span>';
            html += '<span>ВЛИЯНИЕ ' + item.impact + '</span>';
            html += '<span>ДОСТОВЕРНОСТЬ ' + item.confidence + '%</span>';
            html += '<span>📈 ' + item.direction.toUpperCase() + '</span>';
            html += '</div>';
            html += '</div>';
        });
        
        // Обновляем статистику
        document.getElementById('total-signals').textContent = data.length;
        document.getElementById('high-impact').textContent = highImpact;
        document.getElementById('medium-impact').textContent = mediumImpact;
        document.getElementById('avg-confidence').textContent = Math.round(totalConfidence / data.length) + '%';
        document.getElementById('bull-signals').textContent = bullSignals;
        document.getElementById('bear-signals').textContent = bearSignals;
        document.getElementById('active-sectors').textContent = sectors.size;
        document.getElementById('regions').textContent = 1;
        
        document.getElementById('signals-list').innerHTML = html;
        
        // Обновляем время
        const now = new Date();
        const timestamp = now.toLocaleDateString('ru-RU') + ', ' + now.toLocaleTimeString('ru-RU');
        document.getElementById('update-time').textContent = timestamp;
    })
    .catch(error => {
        document.getElementById('signals-list').innerHTML = '<div style="color: #e74c3c; text-align: center; padding: 40px;">Ошибка загрузки данных: ' + error + '</div>';
    });
}

// Автозагрузка при открытии страницы
window.onload = function() {
    setTimeout(loadSignals, 1000);
};

// Автообновление каждые 5 минут
setInterval(loadSignals, 300000);
</script>
</body>
</html>`

	fmt.Fprint(w, html)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	health := map[string]interface{}{
		"status":    "healthy",
		"service":   "SAA Research",
		"version":   "2.0.0",
		"timestamp": time.Now().Format(time.RFC3339),
		"news_count": len(newsData),
	}
	json.NewEncoder(w).Encode(health)
}

func handleNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(newsData)
}

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":  "success",
		"count":   len(newsData),
		"message": fmt.Sprintf("Обновлено %d инвестиционных сигналов", len(newsData)),
	}
	json.NewEncoder(w).Encode(response)
}
