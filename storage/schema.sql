CREATE TABLE IF NOT EXISTS shopping_list_items (
                                                   id INTEGER PRIMARY KEY AUTOINCREMENT,
                                                   item TEXT NOT NULL,
                                                   quantity_owned INTEGER DEFAULT 0,
                                                   quantity_required INTEGER NOT NULL DEFAULT 1,
                                                   is_marked BOOLEAN DEFAULT FALSE,
                                                   last_edit DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                   added_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);