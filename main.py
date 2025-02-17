import os
import curses
import sys

def list_folders():
    """Gibt eine Liste aller Ordner im aktuellen Verzeichnis zurück."""
    return [f for f in os.listdir() if os.path.isdir(f)]

def select_folder(stdscr):
    curses.curs_set(0)
    stdscr.clear()
    
    folders = list_folders()
    if not folders:
        stdscr.addstr(0, 0, "Keine Ordner gefunden.")
        stdscr.refresh()
        stdscr.getch()
        return None
    
    selected = 0
    while True:
        stdscr.clear()
        for i, folder in enumerate(folders):
            if i == selected:
                stdscr.addstr(i, 0, f"> {folder}", curses.A_REVERSE)
            else:
                stdscr.addstr(i, 0, f"  {folder}")
        
        stdscr.refresh()
        key = stdscr.getch()
        
        if key == curses.KEY_UP and selected > 0:
            selected -= 1
        elif key == curses.KEY_DOWN and selected < len(folders) - 1:
            selected += 1
        elif key == 10:  # Enter-Taste
            return folders[selected]

if __name__ == "__main__":
    try:
        folder = curses.wrapper(select_folder)
        if folder:
            print(folder)  # WICHTIG: Go liest diese Ausgabe
    except Exception as e:
        print(f"Fehler: {e}", file=sys.stderr)
        sys.exit(1)
