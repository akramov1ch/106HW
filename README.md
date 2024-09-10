# Uy vazifasi: Fayl Yuklash va Yuklab Olish API bajarishi

## Topshiriq Tavsifi:
`Go` da foydalanuvchilarga fayllarni serverga yuklash va ularni keyinroq yuklab olish imkonini beruvchi oddiy "Fayl Yuklash" va "Yuklab Olish" `API` yarating.

## Endpointlar:
- `POST /upload`: Foydalanuvchilarga fayl yuklashga imkon beradi.
- `GET /download/{filename}`:  Foydalanuvchilarga fayl nomini ko'rsatib faylni yuklab olish imkonini beradi.

## Vazifa Ko'rsatmalari:
1. **Fayl Yuklash:**
    - Fayllarni multipart shakl orqali yuklashni qabul qiling (`multipart/form-data`).
    - Yuklangan fayllarni serverdagi bir katalogga (masalan, `./uploads`) saqlang.
    - Fayl hajmini tekshiring (maksimum `10MB` bilan cheklang).
    - Muvaffaqiyatli javob qaytarib bering, unda fayl nomi va fayl yo'li bo'lishi kerak.

2. **Fayl Yuklab Olish**
    - Foydalanuvchilarga fayl nomini ko'rsatib faylni yuklab olishga imkon bering (masalan, `GET /download/myfile.txt`).
    - Fayl mavjudligini tekshiring; mavjud bo'lmasa, `404 Not Found` statusini qaytaring.
    - Faylni yuklab olish uchun ilova sifatida xizmat qiling.

3. **Tekshiruv va Xatolarni Boshqarish:**
    - Xatolarni to'g'ri boshqaring (masalan, noto'g'ri fayl turlari, fayl hajmi cheklangan, yuklab olinadigan fayl mavjud emas).
    - Faqat ma'lum fayl turlariga ruxsat bering (masalan, `.txt`, `.jpg`, `.png`, `.pdf`).

4. **Xavfsizlikni Ko'rib Chiqish**:
    - Fayl nomlarini tozalab, katalogga kirish hujumlarini oldini oling.
    - Ijro etiladigan fayllarni yuklashga ruxsat bermang (masalan, `.exe`, `.sh`).

5. **Unit Testlari**:
    - Fayl yuklash va yuklab olish funktsiyalari uchun `net/http/httptest` paketidan foydalanib unit testlari yozing.
    - Fayl yuklashni sinovdan o'tkazib, javoblarni tekshiring.

6. **Bonus**:
    - Fayl metama'lumotlarini (masalan, fayl nomi, hajmi, yuklash vaqti) ma'lumotlar bazasida (`PostgreSQL`, yoki `MongoDB`) saqlang.
    - Fayllarning metama'lumotlari bilan birga barcha yuklangan fayllarni ko'rsatuvchi `GET /files` endpointini amalga oshiring.














 















