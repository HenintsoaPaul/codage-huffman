from PIL import Image

def chargerImage(imgPath: str) :
    return Image.open(imgPath).convert('L') # convert to grayscale

def obtenir_intensites_gris(img: Image) :
    width, height = img.size
    intensities = []
    for y in range(height) :
        for x in range(width) :
            intensity = img.getpixel((x, y))
            intensities.append(intensity)
    return intensities

def convertir_en_binaire(intensites):
    binaire = [format(i, '08b') for i in intensites]
    return binaire

if __name__ == "__main__":
    image_path = "test_3x3.png"  # Remplacez par le chemin de votre image
    image = chargerImage(image_path)
    intensities = obtenir_intensites_gris(image)
    
    # Conversion de la première intensité en binaire pour l'exemple
    for index, intensity in enumerate(intensities):
        print(f"{index} -> {intensity}")