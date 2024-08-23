CREATE TABLE IF NOT EXISTS shopping_list_items (
                                                   item TEXT NOT NULL,                            -- nom de l'article
                                                   quantity_owned INTEGER DEFAULT 0,              -- quantité possédée avec valeur par défaut 0
                                                   quantity_required INTEGER NOT NULL DEFAULT 1,  -- quantité requise avec valeur par défaut 1
                                                   is_marked BOOLEAN DEFAULT FALSE,               -- booléen pour marquer l'article, avec valeur par défaut FAUX
                                                   last_edit DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- date/heure de dernière modification
                                                   added_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP   -- date/heure d'ajout de l'article
);
