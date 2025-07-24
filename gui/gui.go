package gui

import (
    "fmt"
    "os/exec"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"

	"github.com/PhateValleyman/copier/config"
	"github.com/PhateValleyman/copier/gui"
)

func Run(entries []config.Entry) {
    a := app.New()
    w := a.NewWindow("Copier")

    selectEntry := widget.NewSelect(func() []string {
        names := []string{}
        for _, e := range entries {
            names = append(names, e.Name)
        }
        return names
    }(), func(s string) {})

    progressBar := widget.NewProgressBar()
    logArea := widget.NewMultiLineEntry()
    logArea.SetReadOnly(true)

    copyBtn := widget.NewButton("Kopírovat", func() {
        idx := selectEntry.Selected
        if idx < 0 {
            dialog.ShowInformation("Chyba", "Vyberte položku", w)
            return
        }
        entry := entries[idx]
        logArea.SetText("")
        go worker.Copy(entry,
            func(p float64) { a.QueueUpdate(func() { progressBar.SetValue(p) }) },
            func(msg string) { a.QueueUpdate(func() { logArea.SetText(logArea.Text + msg + "\n") }) },
        )
        go func() {
            exec.Command("buzzerc", "-t", "5").Run()
        }()
    })

    content := container.NewVBox(
        widget.NewLabel("Vyberte pohádky:"),
        selectEntry,
        copyBtn,
        progressBar,
        widget.NewLabel("Log:"),
        container.NewScroll(logArea),
    )

    w.SetContent(content)
    w.Resize(fyne.NewSize(400, 400))
    w.ShowAndRun()
}
