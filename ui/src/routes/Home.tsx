import { Link } from "react-router-dom";
import { motion } from "framer-motion";
import { BookMarked, ArrowRight, Zap, BarChart2, Link2 } from "lucide-react";
import { buttonVariants } from "@/components/ui/button";

const MotionLink = motion(Link);
const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;



type FeatureCardProps = {
  feature: string;
  icon: React.ComponentType<{ className?: string }>;
  description: string;
  delay?: number;
};

const FeatureCard = ({
  feature,
  icon: Icon,
  description,
  delay = 0,
}: FeatureCardProps) => (
  <motion.div
    className="bg-background/30 backdrop-blur-xl rounded-xl p-6 shadow-lg hover:shadow-xl transition-all duration-300 ease-in-out border border-border"
    whileHover={{
      scale: 1.05,
      rotateY: 15,
    }}
    whileTap={{ scale: 0.95 }}
    initial={{ opacity: 0, y: 20 }}
    animate={{ opacity: 1, y: 0 }}
    transition={{ duration: 0.5, delay }}
  >
    <Icon className="w-12 h-12 mb-4 text-primary-lighter" />
    <h3 className="text-xl font-bold mb-2 text-foreground">{feature}</h3>
    <p className="text-muted-foreground">{description}</p>
  </motion.div>
);

export default function HomePage() {
  return (
    <div className="relative min-h-screen overflow-hidden bg-background text-foreground">
      <div className="absolute inset-0 overflow-hidden">
        <div className="absolute -top-1/4 -left-1/4 w-1/2 h-1/2 bg-primary rounded-full filter blur-[100px] opacity-20 animate-pulse"></div>
        <div className="absolute -bottom-1/4 -right-1/4 w-1/2 h-1/2 bg-primary rounded-full filter blur-[100px] opacity-20 animate-pulse"></div>
        <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-full h-full bg-gradient-to-br from-primary/20 via-primary-dark/10 to-primary-light/20 rounded-full filter blur-[120px] animate-spin-slow"></div>
      </div>

      <div className="relative z-10 mt-5 lg:mt-0">
        <section className="w-full max-w-screen-xl mx-auto min-h-screen flex flex-col justify-center items-center gap-8 pb-8 pt-6 md:py-10 px-4">
          <motion.div
            className="flex max-w-[980px] flex-col items-center gap-4 text-center"
            initial={{ opacity: 0, y: -50 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, ease: "easeOut" }}
          >
            <motion.h1
              className="text-5xl font-extrabold leading-tight tracking-tighter md:text-7xl bg-clip-text text-transparent bg-gradient-to-r from-primary-lighter via-primary-light to-primary-accent"
              initial={{ scale: 0.9, opacity: 0 }}
              animate={{ scale: 1, opacity: 1 }}
              transition={{ duration: 0.8, delay: 0.2 }}
            >
              Discover Your Clusters with Axiom
            </motion.h1>
            <motion.p
              className="max-w-[700px] text-xl text-muted-foreground mt-2 font-normal"
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ duration: 0.8, delay: 0.6 }}
            >
              Axiom gives you visibility into all clusters. Manage everything in
              one place.
            </motion.p>
          </motion.div>

          <motion.div
            className="flex gap-6"
            initial={{ opacity: 0, y: 50 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, delay: 0.8 }}
          >
            <MotionLink
              to="/clusters"
              className={
                buttonVariants({ size: "lg", variant: "default" }) +
                " bg-gradient-to-r from-primary-dark to-primary-lighter text-white font-semibold px-8 py-4 rounded-full"
              }
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
            >
              Start
              <ArrowRight className="ml-2 h-5 w-5" />
            </MotionLink>
            <MotionLink
              to={`https://${BACKEND_URL}/swagger/index.html`}
              target="_blank"
              rel="noreferrer"
              className={
                buttonVariants({ variant: "secondary", size: "lg" }) +
                " border-2 font-semibold px-8 py-4 rounded-full"
              }
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
            >
              <BookMarked className="mr-2 h-5 w-5" />
              <span>Documentation</span>
            </MotionLink>
          </motion.div>

          <motion.div
            className="mt-16 grid grid-cols-1 md:grid-cols-3 gap-8 w-full max-w-5xl"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ duration: 0.8, delay: 1 }}
          >
            <FeatureCard
              feature="Live Metrics"
              icon={BarChart2}
              description="Monitor real-time statistics for every cluster in your system."
              delay={1.2}
            />
            <FeatureCard
              feature="Quick Actions"
              icon={Link2}
              description="Easily access links, shortcuts, and operations related to each environment."
              delay={1.4}
            />
            <FeatureCard
              feature="Streamlined Interactions"
              icon={Zap}
              description="Explore clusters and their details with a lightning-fast interface."
              delay={1.6}
            />
          </motion.div>
        </section>
      </div>
    </div>
  );
}
