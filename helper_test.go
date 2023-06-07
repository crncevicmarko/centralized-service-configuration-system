package main

// import (
// 	"bytes"
// 	"context"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// func TestRenderJSON(t *testing.T) {
// 	ctx := context.Background()
// 	w := httptest.NewRecorder()
// 	v := map[string]string{"key": "value"}

// 	renderJSON(ctx, w, v)

// 	if w.Header().Get("Content-Type") != "application/json" {
// 		t.Errorf("Expected Content-Type header to be 'application/json', got '%s'", w.Header().Get("Content-Type"))
// 	}

// 	if w.Body.String() != `{"key":"value"}` {
// 		t.Errorf("Expected response body to be '{\"key\":\"value\"}', got '%s'", w.Body.String())
// 	}
// }

// У овом примеру користимо два if израза да би проверили да ли заглавље Content-Type и тело одговора имају очекиване вредности. Ако немају, користимо метод t.Errorf да би пријавили грешку. Први аргумент метода t.Errorf је форматни стринг који описује поруку о грешци. Остали аргументи се користе за попуњавање чувара места у форматном стрингу.
// func TestDecodeConfigBody(t *testing.T) {
// 	ctx := context.Background()
// 	r := bytes.NewBufferString(`{
//     "version":"2",
//     "entries": {
//         "key2": "value1",
//         "key3": "value2"
//     }
// }`)

// 	config, err := decodeConfigBody(ctx, r)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "2", config.Version)
// 	assert.Equal(t, "value1", config.Entries["key2"])
// 	assert.Equal(t, "value2", config.Entries["key3"])
// }

// Овај тест креира нови context.Context и bytes.Buffer који садржи дати JSON стринг да би их проследио функцији decodeConfigBody. Затим тест позива функцију decodeConfigBody са овим аргументима и проверава да ли враћа грешку и да ли враћена config вредност има очекивана поља.

// Имајте на уму да овај тест претпоставља да тип cs.Config има поља Version и Entries типа string и map[string]string, респективно. Можда ћете морати да модификујете овај тест да би одговарао стварној дефиницији типа cs.Config у вашем коду.

// func TestCreateId(t *testing.T) {
// 	id := createId()

// 	_, err := uuid.Parse(id)
// 	assert.NoError(t, err)
// }

// Овај тест позива функцију createId да би генерисао нови ИД. Затим користи функцију uuid.Parse из пакета github.com/google/uuid да би парсирао генерисани ИД као UUID. Тест проверава да ли функција uuid.Parse враћа грешку, што указује на то да је генерисани ИД валидан UUID.