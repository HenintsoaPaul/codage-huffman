# from Pillow import Image
from PIL import Image
import numpy as np

# Créer une matrice 3x3 d'intensités de gris prédéfinies
intensites = [
    [0,   32,  64],    # Ligne 1
    [128, 64, 32],   # Ligne 2
    [64, 128, 255]    # Ligne 3
]

# Convertir en tableau numpy et sauvegarder en PNG
img = Image.fromarray(np.array(intensites, dtype=np.uint8), mode='L')
# img.convert("RGB").save("test_3x3.png")
img.save("test_3x3.png")