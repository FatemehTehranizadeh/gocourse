package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

type ctxHandler struct{
	slog.Handler
}

func (h *ctxHandler) Handle(ctx context.Context, r slog.Record) error {
	if city, ok := ctx.Value("city").(string); ok {
		r.AddAttrs(slog.String("city", city))
	} else {
		slog.Error("There is not such a key:", "city")
	}
	return h.Handler.Handle(ctx, r)
}


func main() {

	ctx := context.WithValue(context.Background(), "city", "New York")

	logFile, _ := os.OpenFile("slog_logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	baseHandler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(&ctxHandler{baseHandler})
	slog.SetDefault(logger)
	slog.LogAttrs(ctx,slog.LevelDebug, " ", slog.String("Name", "Reera"))
	slog.With(slog.TimeValue(time.Now()))
	slog.InfoContext(ctx, "First Info Message:")
}




// package main

// import (
// 	"context"
// 	"log/slog"
// 	"os"
// 	"time"
// )

// func main() {

// 	var name string = "Reera"
// 	var id int = 0

// 	ctx := context.WithValue(context.Background(), "city", "New York")

// 	logFile, err := os.OpenFile("slog_loggs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
// 	if err != nil {
// 		slog.Info("Error during opening the file")
// 	}

// 	logger := slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}))
// 	slog.SetDefault(logger)
// 	// slog.Info("This is an info message", "name", name)
// 	// slog.Info("User Info:", slog.String("name", name), slog.Int("ID", id))
// 	slog.Debug("This is an debugging message")
// 	child := logger.WithGroup("request").With(slog.String("method", "Get"))

// 	child.Info("A message with Child")
// 	slog.Info("Third message")
// 	slog.Debug("A debugging message")

// 	// slog.LogAttrs(ctx, slog.LevelDebug, "Debugging message: ", slog.String("name", name), slog.Int("ID", id),
// 	// 	slog.Group("request", slog.String("method", "Get")), slog.Time("Time", time.Now()))

// 	child.LogAttrs(ctx, slog.LevelError, "ERROR: ")

// 	slog.LogAttrs(ctx, slog.LevelDebug, "Debugging message: ", slog.String("name", name), slog.Int("ID", id))

// }
