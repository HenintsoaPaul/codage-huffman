import wave
import struct
import math

def generer_wav(nom_fichier="test.wav", duree=0.0001, freq=440.0, sample_rate=44100, amplitude=32767):
    """Génère un fichier WAV contenant une onde sinusoïdale."""
    num_samples = int(sample_rate * duree)
    
    with wave.open(nom_fichier, "w") as wav_file:
        wav_file.setnchannels(1)  # Mono
        wav_file.setsampwidth(2)  # 16 bits = 2 octets
        wav_file.setframerate(sample_rate)

        for i in range(num_samples):
            # Générer une onde sinusoïdale
            sample = int(amplitude * math.sin(2 * math.pi * freq * i / sample_rate))
            wav_file.writeframes(struct.pack("<h", sample))  # "<h" = 16-bit signé little-endian

    print(f"Fichier {nom_fichier} généré avec succès !")

# Générer un petit fichier test.wav
generer_wav()
